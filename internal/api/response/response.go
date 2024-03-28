package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  string      `json:"Status"`
	Data    interface{} `json:"Data,omitempty"`
	Message interface{} `json:"Message,omitempty"`
}

func SendBadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status:  "fail",
		Message: message,
	})
}

func SendInternalError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Status:  "error",
		Message: message,
	})
}
