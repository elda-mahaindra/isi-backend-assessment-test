package handler

import (
	"context"
	"database/sql"

	"tabungan-api/dto"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) Deposit(c *fiber.Ctx) error {
	var request dto.DepositRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.DepositErrorResponse{
			Remark: "failed to parse request body",
		})
	}

	// call service layer
	saldo, err := handler.service.Deposit(context.Background(), request)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(&dto.DepositErrorResponse{
				Remark: "nomor rekening tidak dikenali`",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(&dto.DepositErrorResponse{
			Remark: "internal server error",
		})
	}

	response := &dto.DepositSuccessResponse{
		Saldo: saldo,
	}

	return c.JSON(response)
}
