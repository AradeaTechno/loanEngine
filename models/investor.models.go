package models

import (
	"amarthaloan/db"
	"amarthaloan/helpers"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func CreateInvestor(objInvestor *helpers.InvestorStruct) (int, error) {
	con := db.CreateConn()
	if err := con.First(&objInvestor, "email = ?", objInvestor.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := con.Create(&objInvestor).Error
			if err != nil {
				return http.StatusInternalServerError, err
			}
		} else {
			return http.StatusInternalServerError, err
		}
	} else {
		return http.StatusConflict, errors.New("investor is exists")
	}
	return http.StatusCreated, nil
}

func IsInvestorExists(investorId string) (bool, error) {
	var count int64
	var objInvestor helpers.InvestorStruct
	con := db.CreateConn()
	if err := con.First(&objInvestor).Where("investor_id = ?", investorId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
