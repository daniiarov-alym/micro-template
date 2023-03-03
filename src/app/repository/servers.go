package repository

import (
	"github.com/daniiarov-alym/micro-template/src/app/service"
	"github.com/jackc/pgx/v4/pgxpool"
	logger "github.com/sirupsen/logrus"
)

type repository struct {
	pool *pgxpool.Pool
}

func NewServerRepository(pool *pgxpool.Pool) service.Repository {
	if pool == nil {
		logger.Fatal("pool cannot be nil")
	}
	return &repository{pool: pool}
}
