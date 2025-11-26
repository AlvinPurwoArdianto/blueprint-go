package model

import "time"

type CategoryRequest struct {
	Id           []int      `json:"id"`
	NameCategory string     `json:"name_category"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type CategoryResponse struct {
	Id           int        `json:"id"`
	NameCategory string     `json:"name_category"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type BulkDeleteCategory struct {
	Id []int `json:"id"`
}
