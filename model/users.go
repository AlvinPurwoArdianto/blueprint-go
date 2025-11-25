package model

import "time"

type UsersResponse struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UsersRequest struct {
	Id        []int      `json:"id"`
	Username  string     `form:"username"`
	Email     string     `form:"email"`
	Password  string     `form:"password"`
	CreatedAt *time.Time `form:"created_at"`
	UpdatedAt time.Time  `form:"updated_at"`
}

type BulkDeleteUsers struct {
	Id []int `json:"id"`
}
