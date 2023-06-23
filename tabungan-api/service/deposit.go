package service

import (
	"context"
	"fmt"

	db "tabungan-api/db/store"
	"tabungan-api/dto"
	"tabungan-api/utils/errs"

	"github.com/sirupsen/logrus"
)

func (service *Service) Deposit(ctx context.Context, request dto.DepositRequest) (saldo int64, err error) {
	const op errs.Op = "service/Deposit"

	// log the request for data tracing purpose
	service.logger.WithFields(logrus.Fields{
		"op": op,
	}).Debug(fmt.Printf("request: %+v", request))

	// call data access layer
	result, err := service.store.DepositTx(ctx, db.DepositTxParams{
		Nominal:    request.Nominal,
		NoRekening: request.NoRekening,
	})
	if err != nil {
		return -1, err
	}

	return result.Account.Saldo, nil
}
