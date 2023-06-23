package service

import (
	"context"
	"fmt"

	db "tabungan-api/db/store"
	"tabungan-api/dto"
	"tabungan-api/utils/errs"

	"github.com/sirupsen/logrus"
)

func (service *Service) Registration(ctx context.Context, request dto.RegistrationRequest) (noRekening string, err error) {
	const op errs.Op = "service/Registration"

	// log the request for data tracing purpose
	service.logger.WithFields(logrus.Fields{
		"op": op,
	}).Debug(fmt.Printf("request: %+v", request))

	// call data access layer
	result, err := service.store.RegistrationTx(ctx, db.RegistrationTxParams{
		Nama: request.Nama,
		Nik:  request.Nik,
		NoHp: request.NoHp,
	})
	if err != nil {
		return "", err
	}

	return result.Account.NoRekening, nil
}
