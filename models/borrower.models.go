package models

import (
	"amarthaloan/db"
	"amarthaloan/helpers"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func CreateBorrower(objBorrower *helpers.BorrowerStruct) (int, error) {
	con := db.CreateConn()
	if err := con.First(&objBorrower, "name = ?", objBorrower.Name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := con.Create(&objBorrower).Error
			if err != nil {
				return http.StatusInternalServerError, err
			}
		} else {
			return http.StatusInternalServerError, err
		}
	} else {
		return http.StatusConflict, errors.New("borrower is exists")
	}
	return http.StatusCreated, nil
}

func IsBorrowerExists(idBorrower string) (bool, error) {
	var count int64
	var objBorrower helpers.BorrowerStruct
	con := db.CreateConn()
	if err := con.First(&objBorrower).Where("borrower_id = ?", idBorrower).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
