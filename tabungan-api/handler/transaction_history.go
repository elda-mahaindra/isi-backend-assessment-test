package handler

import (
	"context"
	"database/sql"

	"tabungan-api/db/sqlc"
	"tabungan-api/dto"
	"tabungan-api/model"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) TransactionHistory(c *fiber.Ctx) error {
	noRekening := c.Params("no_rekening", "")

	request := dto.TransactionHistoryRequest{
		NoRekening: noRekening,
	}

	// call service layer
	entries, err := handler.service.TransactionHistory(context.Background(), request)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(&dto.TransactionHistoryErrorResponse{
				Remark: "nomor rekening tidak dikenali`",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(&dto.TransactionHistoryErrorResponse{
			Remark: "internal server error",
		})
	}

	response := &dto.TransactionHistorySuccessResponse{
		Mutasi: func(entries []sqlc.Entry) []model.Statement {
			mutasi := []model.Statement{}

			for _, entry := range entries {
				mutasi = append(mutasi, model.Statement{
					KodeTransaksi: entry.Code,
					Nominal:       entry.Nominal,
					Waktu:         entry.CreatedAt,
				})
			}

			return mutasi
		}(entries),
	}

	return c.JSON(response)
}
