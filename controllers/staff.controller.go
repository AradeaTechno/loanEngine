package controllers

import (
	"amarthaloan/helpers"
	"amarthaloan/models"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateStaff(c echo.Context) error {
	var objStaff helpers.StaffStruct

	objStaff.Name = c.FormValue("staff_name")
	objStaff.Email = c.FormValue("staff_email")
	objStaff.Role = c.FormValue("staff_role")
	// SANITIZE
	objStaff.Sanitize()
	validate := helpers.ValidateStruct(objStaff)
	if len(validate) > 0 {
		for _, msg := range validate {
			return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New(msg), nil))
		}
	}

	now := time.Now()
	objStaff.CreatedAt = &now
	objStaff.ModifiedAt = &now
	result, err := models.CreateStaff(&objStaff)
	if err != nil {
		return c.JSON(result, helpers.ApiResponse(err, nil))
	}
	return c.JSON(result, helpers.ApiResponse(errors.New("user successfully created"), nil))
}
