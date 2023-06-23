package handler

import (
	"context"
	"database/sql"

	"tabungan-api/dto"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) BalanceCheck(c *fiber.Ctx) error {
	noRekening := c.Params("no_rekening", "")

	request := dto.BalanceCheckRequest{
		NoRekening: noRekening,
	}

	// call service layer
	saldo, err := handler.service.BalanceCheck(context.Background(), request)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(&dto.BalanceCheckErrorResponse{
				Remark: "nomor rekening tidak dikenali`",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(&dto.BalanceCheckErrorResponse{
			Remark: "internal server error",
		})
	}

	response := &dto.BalanceCheckSuccessResponse{
		Saldo: saldo,
	}

	return c.JSON(response)
}
