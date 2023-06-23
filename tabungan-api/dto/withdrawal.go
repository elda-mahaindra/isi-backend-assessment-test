package dto

type WithdrawalRequest struct {
	Nominal    int64  `json:"nominal"`
	NoRekening string `json:"no_rekening"`
}

type WithdrawalErrorResponse struct {
	Remark string `json:"remark"`
}

type WithdrawalSuccessResponse struct {
	Saldo int64 `json:"saldo"`
}
