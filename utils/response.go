package utils

import (
	"strconv"
)

type Meta = map[string]any

// SuccessResponse representasi response sukses
type SuccessResponse[T any] struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Metadata Meta   `json:"metadata,omitempty"`
	Data     T      `json:"data,omitempty"`
}

// ErrorResponse representasi response error
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SuccessResponseNoData(message string) SuccessResponse[any] {
	return SuccessResponse[any]{
		Success: true,
		Message: message,
	}
}

// Ok buat response sukses
func Ok[T any](data T, message string, metadata Meta) SuccessResponse[T] {
	if message == "" {
		message = "success"
	}
	return SuccessResponse[T]{
		Success:  true,
		Message:  message,
		Metadata: metadata,
		Data:     data,
	}
}

// Fail buat response error
func Fail(message string, metadata interface{}) ErrorResponse {
	if message == "" {
		message = "failed"
	}
	return ErrorResponse{
		Success: false,
		Message: message,
	}
}

func OkPaginate[T any](items []T, meta Pagination, message string) SuccessResponse[[]T] {
	return Ok(items, message, Meta{
		"page":        meta.Page,
		"per_page":    meta.PerPage,
		"total":       meta.Total,
		"total_pages": meta.TotalPages,
	})
}

// GetPaginationParams ambil page & per_page dari input query
func GetPaginationParams(query map[string]string, defaults map[string]int) (int, int) {
	perPage := defaults["per_page"]
	page := defaults["page"]

	if val, ok := query["per_page"]; ok {
		if n := toInt(val); n > 0 {
			perPage = n
		}
	}
	if val, ok := query["page"]; ok {
		if n := toInt(val); n > 0 {
			page = n
		}
	}

	if perPage <= 0 {
		perPage = 10
	}
	if page <= 0 {
		page = 1
	}

	return perPage, page
}

// helper konversi string -> int

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}
