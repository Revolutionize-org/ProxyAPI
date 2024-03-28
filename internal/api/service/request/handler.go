package request

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"gitlab.com/revolutionize1/foward-api/internal/api/response"
	"gitlab.com/revolutionize1/foward-api/internal/api/validation"
)

func ReceiveRequest(c *fiber.Ctx) error {
	requestInfo := new(requestInformation)

	if err := c.BodyParser(requestInfo); err != nil {
		return response.SendBadRequest(c, "no payload was provided")
	}

	validationErr := validation.Validate(requestInfo)

	if validationErr != nil {
		errors := formatValidationErrors(validationErr)
		return c.Status(fiber.StatusBadRequest).JSON(response.Response{
			Status: "fail",
			Data:   errors,
		})
	}

	apiKey := c.Get("X-API-KEY")
	if requestInfo.Header == nil {
		return response.SendBadRequest(c, "request nil")
	}
	requestInfo.Header["X-API-KEY"] = apiKey

	resp, err := selectRequestMethod(requestInfo)

	if err != nil {
		if err.Error() == "invalid method provided" {
			return response.SendBadRequest(c, err.Error())
		}
		return response.SendInternalError(c, err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.SendInternalError(c, err.Error())
	}

	responseData := createSuccessResponse(resp, string(body))
	return c.JSON(responseData)
}

func selectRequestMethod(requestInfo *requestInformation) (*http.Response, error) {
	switch requestInfo.Method {
	case "GET", "POST", "PUT", "DELETE", "PATCH":
		return query(requestInfo)
	default:
		return nil, errors.New("invalid method provided")
	}
}

func formatValidationErrors(errs []validation.ValidationError) []map[string]string {
	var errors []map[string]string

	for _, err := range errs {
		errors = append(errors, err.Format())
	}
	return errors
}

func createSuccessResponse(resp *http.Response, body string) *response.Response {
	return &response.Response{
		Status: "success",
		Data: &requestResponse{
			Status:        resp.Status,
			StatusCode:    resp.StatusCode,
			Proto:         resp.Proto,
			ProtoMajor:    resp.ProtoMajor,
			ProtoMinor:    resp.ProtoMinor,
			Header:        resp.Header,
			Body:          json.RawMessage(body),
			ContentLength: resp.ContentLength,
		},
	}
}
