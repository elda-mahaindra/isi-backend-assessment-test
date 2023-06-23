package dto

type BalanceCheckRequest struct {
	NoRekening string `json:"no_rekening"`
}

type BalanceCheckErrorResponse struct {
	Remark string `json:"remark"`
}

type BalanceCheckSuccessResponse struct {
	Saldo int64 `json:"saldo"`
}
