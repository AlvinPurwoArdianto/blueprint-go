package model

import "time"

type UsersRequest struct {
	Id        []int      `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type UsersResponse struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type BulkDeleteUsers struct {
	Id []int `json:"id"`
}
