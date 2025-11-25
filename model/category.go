package model

import "time"

type CategoryResponse struct {
	Id           int        `json:"id"`
	NameCategory string     `json:"name_category"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type CategoryRequest struct {
	Id           []int      `json:"id"`
	NameCategory string     `form:"name_category"`
	CreatedAt    *time.Time `form:"created_at"`
	UpdatedAt    time.Time  `form:"updated_at"`
}

type BulkDeleteCategory struct {
	Id []int `json:"id"`
}
