package service

import (
	"confluence-payment/core-internal/infrastructure"
	core_model "confluence-payment/core-internal/model"
	"confluence-payment/core-internal/utils"
	"confluence-payment/feature-payment/v1/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PaymentService struct {
	PostgresInfra           *infrastructure.PostgresInfra
	RedisInfra              *infrastructure.RedisInfra
	OtelInfra               *infrastructure.OtelInfra
}

func NewPaymentService(postgresInfra *infrastructure.PostgresInfra, redisInfra *infrastructure.RedisInfra, otelInfra *infrastructure.OtelInfra) *PaymentService {
	return &PaymentService{
		PostgresInfra:           postgresInfra,
		RedisInfra:              redisInfra,
		OtelInfra:               otelInfra,
	}
}

func (p *PaymentService) CreatePayment(ctx context.Context, data model.Payment) (res core_model.CoreIdResponse, traceID string, statusCode int, err error) {

	paymentUUID := uuid.New().String()
	data.UUID = paymentUUID

	paymentInterface := [][]interface{}{
		{data.UUID, data.OrderID, data.ShopID, data.CustomerID, data.Amount},
	}

	paymentString := utils.PrepareInsertQuery(utils.PaymentTableName, utils.PaymentColumnsListForInsert, paymentInterface)

	var paymentParams []interface{}
	for _, payment := range paymentInterface {
		paymentParams = append(paymentParams, payment...)
	}

	

	tx, err := p.PostgresInfra.DbWritePool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return res, traceID, http.StatusInternalServerError, err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, paymentString, paymentParams...)
	if err != nil {
		return res, traceID, http.StatusInternalServerError, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return res, traceID, http.StatusInternalServerError, err
	}

	res.ID = paymentUUID
	
	return res, traceID, http.StatusOK, nil
}

func (p *PaymentService) DetailPayment(ctx context.Context, id string, data model.Payment) (res model.Payment, traceID string, statusCode int, err error) {
	ctx, span := p.OtelInfra.Tracer.Start(ctx, "DetailPayment")
	traceID = span.SpanContext().TraceID().String()

	paymentString := fmt.Sprintf(
		`SELECT id, created_at, updated_at, deleted_at, uuid, order_id, customer_id, amount FROM %s WHERE deleted_at IS NULL AND uuid = $1;`, utils.PaymentTableName,
	)

	cache, err := p.RedisInfra.Client.Get(paymentString).Bytes()
	if err == nil {
		err = json.Unmarshal(cache, &res)
		if err == nil {
			return res, traceID, http.StatusOK, nil
		}
	}

	rows, err := p.PostgresInfra.DbReadPool.Query(ctx, paymentString, id)
	if err != nil {
		span.End()
		return res, traceID, http.StatusInternalServerError, err
	}
	defer rows.Close()


	payment := model.Payment{}

	for rows.Next() {
		var rawPayment model.Payment

		err = rows.Scan(
			&rawPayment.ID,
			&rawPayment.CreatedAt,
			&rawPayment.UpdatedAt,
			&rawPayment.DeletedAt,
			&rawPayment.UUID,
			&rawPayment.OrderID,
			&rawPayment.CustomerID,
			&rawPayment.Amount,
		)
		if err != nil {
			span.End()
			return res, traceID, http.StatusInternalServerError, err
		}

		if payment.UUID == "" {
			payment = rawPayment
		}
	}

	p.RedisInfra.Client.Set(paymentString, payment, utils.DefaultRedisTimeOut)

	span.End()

	res = payment
	return res, traceID, http.StatusOK, nil
}

func (p *PaymentService) ListPayment(ctx context.Context, query core_model.CoreQuery) (res []model.PaymentResponse, pagination core_model.Pagination, traceID string, statusCode int, err error) {
	ctx, span := p.OtelInfra.Tracer.Start(ctx, "ListPayment")
	traceID = span.SpanContext().TraceID().String()

	pagination.Page = query.Page
	countString := utils.PrepareSelectCountQuery(utils.PaymentTableName)

	intVariable := 0
	paymentInterface := []interface{}{}
	countInterface := []interface{}{}
	paymentString := utils.PrepareSelectQuery(utils.PaymentTableName, utils.PaymentColumnsListForSelect)

	if query.ShopID != "" {
		intVariable++
		paymentString += fmt.Sprintf(" AND shop_id = $%d", intVariable)
		paymentInterface = append(paymentInterface, query.ShopID)

		countString += fmt.Sprintf(" AND shop_id = $%d", intVariable)
		countInterface = append(countInterface, query.ShopID)
	} else {
		span.End()
		return res, pagination, traceID, http.StatusBadRequest, err
	}

	if query.Sort == "" {
		query.Sort = utils.DefaultOrder
	}
	paymentString += fmt.Sprintf(" ORDER BY %s ", query.Sort)

	intVariable++
	paymentString += fmt.Sprintf(`LIMIT $%d `, intVariable)
	if query.Size == "" {
		query.Size = utils.DefaultSize
	}
	paymentInterface = append(paymentInterface, query.Size)
	pagination.Size = query.Size

	intVariable++
	paymentString += fmt.Sprintf(`OFFSET $%d `, intVariable)
	if query.Page == "" {
		query.Page = utils.DefaultPage
	}
	offset := (utils.StringToInt(query.Page) - 1) * utils.StringToInt(query.Size)
	paymentInterface = append(paymentInterface, utils.IntToString(offset))
	pagination.Page = query.Page

	cache, err := p.RedisInfra.Client.Get(paymentString).Bytes()
	if err == nil {
		err = json.Unmarshal(cache, &res)
		if err == nil {
			return res, pagination, traceID, http.StatusOK, nil
		}
	}

	err = p.PostgresInfra.DbReadPool.QueryRow(ctx, countString, countInterface...).Scan(&pagination.TotalItems)
	if err != nil {
		span.End()
		return res, pagination, traceID, http.StatusInternalServerError, err
	}

	paymentRows, err := p.PostgresInfra.DbReadPool.Query(ctx, paymentString, paymentInterface...)
	if err != nil {
		span.End()
		return res, pagination, traceID, http.StatusInternalServerError, err
	}
	defer paymentRows.Close()

	paymentIDs := []string{}
	payments := []model.PaymentResponse{}
	for paymentRows.Next() {
		var rawPayment model.Payment
		err = paymentRows.Scan(
			&rawPayment.ID,
			&rawPayment.CreatedAt,
			&rawPayment.UpdatedAt,
			&rawPayment.DeletedAt,
			&rawPayment.UUID,
			&rawPayment.OrderID,
			&rawPayment.ShopID,
			&rawPayment.CustomerID,
			&rawPayment.Amount,
		)

		if err != nil {
			return res, pagination, traceID, http.StatusInternalServerError, err
		}

		paymentIDs = append(paymentIDs, fmt.Sprintf(`'%s'`, rawPayment.UUID))
		payments = append(payments, model.PaymentResponse{
			Payment: rawPayment,
		})
	}

	if len(paymentIDs) == 0 {
		span.End()
		return res, pagination, traceID, http.StatusOK, nil
	}

	p.RedisInfra.Client.Set(paymentString, payments, utils.DefaultRedisTimeOut)

	span.End()
	return payments, pagination, traceID, http.StatusOK, nil
}
