package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"todos/controller"
	"todos/db"
)

func Init() error {
	e := echo.New()

	db, err := db.Init()
	if err != nil {
		return err
	}
	defer db.Close()

	e.GET("", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "Application is Running",
		})
	})

	// ROUTE CATEGORY
	c := e.Group("/category")
	c.GET("", controller.GetAllCategory(db))
	c.GET("/:id", controller.GetCategoryById(db))
	c.POST("/create", controller.CreateCategory(db))
	c.PUT("/edit/:id", controller.EditCategory(db))
	c.DELETE("/:id", controller.DeleteCategory(db))
	c.DELETE("", controller.BulkDeleteCategory(db))

	// ROUTE USERS
	u := e.Group("/user")
	u.GET("", controller.GetAllUsers(db))
	u.GET("/:id", controller.GetUsersById(db))
	u.POST("/create", controller.CreateUsers(db))
	u.PUT("/edit/:id", controller.EditUsers(db))
	// u.DELETE("/:id", controller.DeleteUsers(db))
	// u.DELETE("", controller.BulkDeleteUsers(db))

	return e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
