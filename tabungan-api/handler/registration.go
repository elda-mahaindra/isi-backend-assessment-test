package handler

import (
	"context"

	"tabungan-api/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func (handler *Handler) Registration(c *fiber.Ctx) error {
	var request dto.RegistrationRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.RegistrationErrorResponse{
			Remark: "failed to parse request body",
		})
	}

	// call service layer
	noRekening, err := handler.service.Registration(context.Background(), request)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return c.Status(fiber.StatusBadRequest).JSON(&dto.RegistrationErrorResponse{
					Remark: "nik atau nomor hp sudah digunakan",
				})
			}
		}

		return c.Status(fiber.StatusInternalServerError).JSON(&dto.RegistrationErrorResponse{
			Remark: "internal server error",
		})
	}

	response := &dto.RegistrationSuccessResponse{
		NoRekening: noRekening,
	}

	return c.JSON(response)
}
