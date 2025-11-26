package controller

import (
	"database/sql"
	"net/http"
	"todos/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []model.UsersResponse

		query := `
		SELECT id, username, email, password, created_at, updated_at 
		FROM users
		`

		rows, err := db.Query(query)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Terjadi Kesalahan",
				"error":   err.Error(),
			})
		}

		for rows.Next() {
			var user model.UsersResponse
			err := rows.Scan(
				&user.Id,
				&user.Username,
				&user.Email,
				&user.Password,
				&user.CreatedAt,
				&user.UpdatedAt,
			)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "Terjadi Kesalahan",
					"error":   err.Error(),
				})
			}
			users = append(users, user)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Sukses Menampilan Semua data Users",
			"data":    users,
		})
	}
}

func GetUsersById(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users model.UsersResponse
		id := c.Param("id")

		query := `
		SELECT id, username, email, password, created_at, updated_at
			FROM users
		WHERE id = $1
		`

		row := db.QueryRow(query, id)

		err := row.Scan(
			&users.Id,
			&users.Username,
			&users.Email,
			&users.Password,
			&users.CreatedAt,
			&users.UpdatedAt,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "Data Tidak Ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Terjadi Kesalahan",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Sukses Menampilkan data Users Berdasarkan Id",
			"data":    users,
		})
	}
}

func CreateUsers(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users model.UsersResponse
		var req model.UsersRequest

		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid request payload",
				"error":   err.Error(),
			})
		}

		query := `
			INSERT INTO users (username, email, password, created_at)
				VALUES
			($1, $2, $3, now())
				RETURNING
			id, username, email, password, created_at, updated_at
		`

		rows := db.QueryRowx(query, req.Username, req.Email, req.Password)
		err = rows.Scan(
			&users.Id,
			&users.Username,
			&users.Email,
			&users.Password,
			&users.CreatedAt,
			&users.UpdatedAt,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Terjadi Kesalahan",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Berhasil Menambahkan Data Users",
			"data":    users,
		})
	}
}

func EditUsers(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.UsersResponse
		var req model.UsersRequest

		id := c.Param("id")
		err := c.Bind(&req)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid request payload",
				"error":   err.Error(),
			})
		}

		query := `
			UPDATE users SET username = $1, email = $2, password = $3, updated_at = now()
				WHERE id = $4
			RETURNING
				id, username, email, password, created_at, updated_at
		`

		rows := db.QueryRow(query, req.Username, req.Email, req.Password, id)
		err = rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Terjadi Kesalahan",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Berhasil Mengubah Data Users",
			"data":    user,
		})
	}
}
