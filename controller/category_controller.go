package controller

import (
	"database/sql"
	"net/http"
	"todos/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func GetAllCategory(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var categorys []model.CategoryResponse

		query := `
			SELECT id, name_category, created_at, updated_at
			FROM category
		`

		rows, err := db.Query(query)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Terjadi Kesalahan",
				"error":   err.Error(),
			})
		}

		for rows.Next() {
			var category model.CategoryResponse
			err = rows.Scan(
				&category.Id,
				&category.NameCategory,
				&category.CreatedAt,
				&category.UpdatedAt,
			)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "Terjadi Kesalahan",
					"error":   err.Error(),
				})
			}
			categorys = append(categorys, category)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Sukses Menampilan Semua data Category",
			"data":    categorys,
		})
	}
}

func GetCategoryById(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var categorys model.CategoryResponse
		id := c.Param("id")

		query := `
		SELECT id, name_category, created_at, updated_at 
			FROM category 
		WHERE id = $1
		`

		err := db.QueryRow(query, id).Scan(
			&categorys.Id,
			&categorys.NameCategory,
			&categorys.CreatedAt,
			&categorys.UpdatedAt,
		)

		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "Data Tidak Ditemukan",
			})
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Terjadi Kesalahan",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Sukses Menampilkan data Category Berdasarkan Id",
			"data":    categorys,
		})
	}
}

func CreateCategory(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var categorys model.CategoryResponse
		var req model.CategoryRequest

		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid request payload",
				"error":   err.Error(),
			})
		}

		query := `
			INSERT INTO category (name_category, created_at)
				VALUES
			($1, now())
				RETURNING
			id, name_category, created_at, updated_at
		`

		rows := db.QueryRowx(query, req.NameCategory)
		err = rows.Scan(
			&categorys.Id,
			&categorys.NameCategory,
			&categorys.CreatedAt,
			&categorys.UpdatedAt,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Terjadi Kesalahan",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Sukses Menambahkan data Category",
			"data":    categorys,
		})
	}
}

func EditCategory(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.CategoryRequest
		var category model.CategoryResponse

		Id := c.Param("id")
		err := c.Bind(&req)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid request payload",
				"error":   err.Error(),
			})
		}

		query := `
			UPDATE category SET name_category = $1, updated_at = now()
				WHERE id = $2
			RETURNING
				id, name_category, created_at, updated_at
		`

		rows := db.QueryRow(query, req.NameCategory, Id)
		err = rows.Scan(
			&category.Id,
			&category.NameCategory,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Terjadi Kesalahan",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Kategori Berhasil Di Ubah",
			"data":    category,
		})
	}
}

func DeleteCategory(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.CategoryRequest
		Id := c.Param("id")
		err := c.Bind(&req)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid request payload",
				"error":   err.Error(),
			})
		}

		query := `
			DELETE FROM category
			WHERE id = $1
		`

		_, err = db.Exec(query, Id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Data Category Tidak Ditemukan",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Berhasil Mengahapus Data Category",
		})
	}
}

func BulkDeleteCategory(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.BulkDeleteCategory
		err := c.Bind(&req)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid request payload",
				"error":   err.Error(),
			})
		}

		for _, Id := range req.Id {
			query := `
				DELETE FROM category WHERE id = $1
			`
			_, err = db.Exec(query, Id)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "Terjadi Kesalahan",
					"error":   err.Error(),
				})
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Berhasil Menghapus data category berdasarkan id yang sudah dipilih",
		})
	}
}
