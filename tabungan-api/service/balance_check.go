package service

import (
	"context"
	"fmt"

	"tabungan-api/dto"
	"tabungan-api/utils/errs"

	"github.com/sirupsen/logrus"
)

func (service *Service) BalanceCheck(ctx context.Context, request dto.BalanceCheckRequest) (saldo int64, err error) {
	const op errs.Op = "service/BalanceCheck"

	// log the request for data tracing purpose
	service.logger.WithFields(logrus.Fields{
		"op": op,
	}).Debug(fmt.Printf("request: %+v", request))

	// call data access layer
	account, err := service.store.GetAccount(ctx, request.NoRekening)
	if err != nil {
		return -1, err
	}

	return account.Saldo, nil
}
