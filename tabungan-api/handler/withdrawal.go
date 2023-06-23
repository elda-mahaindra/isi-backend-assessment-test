package handler

import (
	"context"
	"database/sql"

	"tabungan-api/dto"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) Withdrawal(c *fiber.Ctx) error {
	var request dto.WithdrawalRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.WithdrawalErrorResponse{
			Remark: "failed to parse request body",
		})
	}

	// call service layer
	saldo, err := handler.service.Withdrawal(context.Background(), request)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(&dto.WithdrawalErrorResponse{
				Remark: "nomor rekening tidak dikenali`",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(&dto.WithdrawalErrorResponse{
			Remark: "internal server error",
		})
	}

	response := &dto.WithdrawalSuccessResponse{
		Saldo: saldo,
	}

	return c.JSON(response)
}
