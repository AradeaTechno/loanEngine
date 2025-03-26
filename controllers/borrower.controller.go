package controllers

import (
	"amarthaloan/helpers"
	"amarthaloan/models"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateBorrower(c echo.Context) error {
	var objBorrower helpers.BorrowerStruct

	objBorrower.Name = c.FormValue("borrower_name")
	objBorrower.Sanitize()
	validate := helpers.ValidateStruct(objBorrower)
	if len(validate) > 0 {
		for _, msg := range validate {
			return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New(msg), nil))
		}
	}

	now := time.Now()
	objBorrower.CreatedAt = &now
	objBorrower.ModifiedAt = &now
	result, err := models.CreateBorrower(&objBorrower)
	if err != nil {
		return c.JSON(result, helpers.ApiResponse(err, nil))
	}
	return c.JSON(http.StatusOK, helpers.ApiResponse(errors.New("borrower successfully created"), nil))
}
