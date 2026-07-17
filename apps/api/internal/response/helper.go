package response

import "github.com/gofiber/fiber/v3"

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}

type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Error      any         `json:"error,omitempty"`
}

func Success(c fiber.Ctx, status int, message string, data any, pagination ...*Pagination) error {
	var paginationData *Pagination
	if len(pagination) > 0 && pagination[0] != nil {
		paginationData = pagination[0]
	}

	return c.Status(status).JSON(Response{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: paginationData,
	})
}

func Error(c fiber.Ctx, status int, message string, err any) error {
	return c.Status(status).JSON(Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}
