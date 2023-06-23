package service

import (
	"context"
	"fmt"

	"tabungan-api/db/sqlc"
	"tabungan-api/dto"
	"tabungan-api/utils/errs"

	"github.com/sirupsen/logrus"
)

func (service *Service) TransactionHistory(ctx context.Context, request dto.TransactionHistoryRequest) (entries []sqlc.Entry, err error) {
	const op errs.Op = "service/BalanceCheck"

	// log the request for data tracing purpose
	service.logger.WithFields(logrus.Fields{
		"op": op,
	}).Debug(fmt.Printf("request: %+v", request))

	// call data access layer
	entries, err = service.store.GetEntries(ctx, request.NoRekening)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
