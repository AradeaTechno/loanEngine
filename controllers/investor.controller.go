package controllers

import (
	"amarthaloan/helpers"
	"amarthaloan/models"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateInvestor(c echo.Context) error {
	var objInvestor helpers.InvestorStruct

	objInvestor.Name = c.FormValue("investor_name")
	objInvestor.Email = c.FormValue("investor_email")
	objInvestor.Sanitize()
	validate := helpers.ValidateStruct(objInvestor)
	if len(validate) > 0 {
		for _, msg := range validate {
			return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New(msg), nil))
		}
	}
	now := time.Now()
	objInvestor.CreatedAt = &now
	objInvestor.ModifiedAt = &now
	result, err := models.CreateInvestor(&objInvestor)
	if err != nil {
		return c.JSON(result, helpers.ApiResponse(err, nil))
	}
	return c.JSON(result, helpers.ApiResponse(errors.New("investor created successfully"), nil))
}
