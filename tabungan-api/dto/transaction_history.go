package dto

type TransactionHistoryRequest struct {
	NoRekening string `json:"no_rekening"`
}

type TransactionHistoryErrorResponse struct {
	Remark string `json:"remark"`
}

type TransactionHistorySuccessResponse struct {
	Mutasi []struct {
		KodeTransaksi string `json:"kode_transaksi"`
		Nominal       int64  `json:"nominal"`
		Waktu         string `json:"waktu"`
	} `json:"mutasi"`
}
