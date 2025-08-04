package service

import (
	"confluence-payment/core-internal/infrastructure"
	payment_proto "confluence-payment/core-internal/protos/payment"
	"confluence-payment/feature-payment/v1/model"
	"context"
)

type PaymentProtoService struct {
	PaymentService *PaymentService
	PostgresInfra  *infrastructure.PostgresInfra
	payment_proto.UnimplementedPaymentProtoServer
}

func NewPaymentProtoService(paymentService *PaymentService, postgresInfra *infrastructure.PostgresInfra) *PaymentProtoService {
	return &PaymentProtoService{
		PaymentService: paymentService,
		PostgresInfra:  postgresInfra,
	}
}

func (p *PaymentProtoService) CreatePayment(ctx context.Context, in *payment_proto.CreatePaymentRequest) (*payment_proto.CreatePaymentResponse, error) {

	payment := model.Payment{
		OrderID:    in.OrderID,
		ShopID:     in.ShopID,
		CustomerID: in.CustomerID,
		Amount:     in.Amount,
	}

	data, _, _, err := p.PaymentService.CreatePayment(ctx, payment)
	if err != nil {
		return nil, err
	}

	if data.ID == "" {
		return nil, err
	}

	// data := model.Payment{}

	// paymentUUID := uuid.New().String()
	// data.UUID = paymentUUID

	// paymentInterface := [][]interface{}{
	// 	{data.UUID, data.OrderID, data.ShopID, data.CustomerID, data.Amount},
	// }

	// paymentString := utils.PrepareInsertQuery(utils.PaymentTableName, utils.PaymentColumnsListForInsert, paymentInterface)

	// var paymentParams []interface{}
	// for _, payment := range paymentInterface {
	// 	paymentParams = append(paymentParams, payment...)
	// }

	// tx, err := p.PostgresInfra.DbWritePool.BeginTx(ctx, pgx.TxOptions{})
	// if err != nil {
	// 	return nil, err
	// }
	// defer tx.Rollback(ctx)

	// _, err = tx.Exec(ctx, paymentString, paymentParams...)
	// if err != nil {
	// 	return nil, err
	// }

	// err = tx.Commit(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	res := &payment_proto.CreatePaymentResponse{}
	res.Success = true

	return res, nil
}
