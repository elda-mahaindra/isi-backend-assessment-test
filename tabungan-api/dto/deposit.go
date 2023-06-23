package dto

type DepositRequest struct {
	Nominal    int64  `json:"nominal"`
	NoRekening string `json:"no_rekening"`
}

type DepositErrorResponse struct {
	Remark string `json:"remark"`
}

type DepositSuccessResponse struct {
	Saldo int64 `json:"saldo"`
}
