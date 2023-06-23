package model

import (
	"time"
)

type Statement struct {
	KodeTransaksi string    `json:"kode_transaksi"`
	Nominal       int64     `json:"nominal"`
	Waktu         time.Time `json:"waktu"`
}
