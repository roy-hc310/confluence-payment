package infrastructure

import (
	"confluence-payment/core-internal/utils"
	"confluence-payment/sqlc"
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresInfra struct {
	DbReadPool  *pgxpool.Pool
	DbWritePool *pgxpool.Pool

	ReadQuery  *sqlc.Queries
	WriteQuery *sqlc.Queries
}

func NewPostgresInfra() *PostgresInfra {
	dsnRead := fmt.Sprintf("host=%s user=%s password=%s dbname=%s search_path=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		utils.GlobalEnv.DBReadHost, utils.GlobalEnv.DBReadUser, utils.GlobalEnv.DBReadPass, utils.GlobalEnv.DBReadName, utils.GlobalEnv.DBReadSchema, utils.GlobalEnv.DBReadPort)

	dbReadConfig, err := pgxpool.ParseConfig(dsnRead)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	dbReadConfig.MaxConns = 10
	dbReadConfig.MinConns = 0
	dbReadConfig.MaxConnLifetime = time.Hour
	dbReadConfig.MaxConnIdleTime = time.Minute * 30
	dbReadConfig.HealthCheckPeriod = time.Minute * 5

	dbReadPool, err := pgxpool.NewWithConfig(context.Background(), dbReadConfig)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	readQuery := sqlc.New(dbReadPool)

	dsnWrite := fmt.Sprintf("host=%s user=%s password=%s dbname=%s search_path=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		utils.GlobalEnv.DBWriteHost, utils.GlobalEnv.DBWriteUser, utils.GlobalEnv.DBWritePass, utils.GlobalEnv.DBWriteName, utils.GlobalEnv.DBWriteSchema, utils.GlobalEnv.DBWritePort)

	dbWriteConfig, err := pgxpool.ParseConfig(dsnWrite)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	dbWriteConfig.MaxConns = 10
	dbWriteConfig.MinConns = 0
	dbWriteConfig.MaxConnLifetime = time.Hour
	dbWriteConfig.MaxConnIdleTime = time.Minute * 30
	dbWriteConfig.HealthCheckPeriod = time.Minute * 5

	dbWritePool, err := pgxpool.NewWithConfig(context.Background(), dbWriteConfig)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	writeQuery := sqlc.New(dbWritePool)

	return &PostgresInfra{
		DbReadPool:  dbReadPool,
		DbWritePool: dbWritePool,

		ReadQuery:  readQuery,
		WriteQuery: writeQuery,
	}
}
