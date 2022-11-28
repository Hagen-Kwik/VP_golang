package controllers

import (
	"net/http"
	"vp_week11_echo/models"

	"github.com/labstack/echo/v4"
)

func FetchAllMahasiswa(c echo.Context) error {

	result, err := models.FetchAllMahasiswa()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreMahasiswa(c echo.Context) error {

	nim := c.FormValue("nim")
	name := c.FormValue("name")
	gender := c.FormValue("gender")
	fakultas := c.FormValue("fakultas")
	prodi := c.FormValue("prodi")

	result, err := models.StoreMahasiswa(nim, name, gender, fakultas, prodi)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateMahasiswa(c echo.Context) error {

	nim := c.FormValue("nim")
	name := c.FormValue("name")
	gender := c.FormValue("gender")
	fakultas := c.FormValue("fakultas")
	prodi := c.FormValue("prodi")

	result, err := models.UpdateMahasiswa(nim, name, gender, fakultas, prodi)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteMahasiswa(c echo.Context) error {

	nim := c.FormValue("nim")

	result, err := models.DeleteMahasiswa(nim)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
