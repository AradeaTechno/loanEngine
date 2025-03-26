package models

import (
	"amarthaloan/db"
	"amarthaloan/helpers"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func CreateStaff(objStaff *helpers.StaffStruct) (int, error) {
	con := db.CreateConn()
	if err := con.First(&objStaff, "email = ?", objStaff.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := con.Create(&objStaff).Error
			if err != nil {
				return http.StatusInternalServerError, err
			}
		} else {
			return http.StatusInternalServerError, err
		}
	} else {
		return http.StatusConflict, errors.New("user is exists")
	}
	return http.StatusCreated, nil
}

func IsStaffExists(idStaff string) (bool, error) {
	var count int64
	var objStaff helpers.StaffStruct
	con := db.CreateConn()
	if err := con.Model(&objStaff).Where("staff_id = ?", idStaff).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
