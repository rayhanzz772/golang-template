package utils

import "math"

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func NewPagination(page, perPage, total int) Pagination {
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	return Pagination{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
	}
}
