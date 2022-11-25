package routes

import (
	"net/http"
	"vp_week11_echo/controllers"

	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello, "+name)
}

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello routes!")
	})

	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello routes/user")
	})

	e.GET("user/:name", getUser)

	e.GET("/mahasiswa", controllers.FetchAllMahasiswa)

	return e
}
