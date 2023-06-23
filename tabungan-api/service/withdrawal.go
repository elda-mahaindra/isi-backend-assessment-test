package service

import (
	"context"
	"fmt"

	db "tabungan-api/db/store"
	"tabungan-api/dto"
	"tabungan-api/utils/errs"

	"github.com/sirupsen/logrus"
)

func (service *Service) Withdrawal(ctx context.Context, request dto.WithdrawalRequest) (saldo int64, err error) {
	const op errs.Op = "service/Withdrawal"

	// log the request for data tracing purpose
	service.logger.WithFields(logrus.Fields{
		"op": op,
	}).Debug(fmt.Printf("request: %+v", request))

	// call data access layer
	result, err := service.store.WithdrawalTx(ctx, db.WithdrawalTxParams{
		Nominal:    request.Nominal,
		NoRekening: request.NoRekening,
	})
	if err != nil {
		return -1, err
	}

	return result.Account.Saldo, nil
}
