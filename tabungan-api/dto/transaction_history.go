package dto

import (
	"tabungan-api/model"
)

type TransactionHistoryRequest struct {
	NoRekening string `json:"no_rekening"`
}

type TransactionHistoryErrorResponse struct {
	Remark string `json:"remark"`
}

type TransactionHistorySuccessResponse struct {
	Mutasi []model.Statement `json:"mutasi"`
}
