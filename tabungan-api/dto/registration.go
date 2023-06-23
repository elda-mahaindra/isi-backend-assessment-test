package dto

type RegistrationRequest struct {
	Nama string `json:"nama"`
	Nik  string `json:"nik"`
	NoHp string `json:"no_hp"`
}

type RegistrationErrorResponse struct {
	Remark string `json:"remark"`
}

type RegistrationSuccessResponse struct {
	NoRekening string `json:"no_rekening"`
}
