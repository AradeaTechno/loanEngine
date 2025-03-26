package models

import (
	"amarthaloan/db"
	"amarthaloan/helpers"
	"net/http"
)

func Disburse(objDisburse *helpers.DisburseStruct) (int, error) {

	con := db.CreateConn()
	tx := con.Begin()
	if tx.Error != nil {
		return http.StatusInternalServerError, tx.Error
	}

	// SAVE DISBURSE
	if err := con.Create(&objDisburse).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}

	var objLoan helpers.LoanStruct
	disburse := map[string]any{
		"state":             "disburse",
		"disbursed_by":      objDisburse.StaffId,
		"disbursement_date": objDisburse.CreatedAt,
		"modified_at":       objDisburse.CreatedAt,
	}
	// UPDATE LOAN DATA
	if err := tx.Model(&objLoan).Where("loan_id = ?", objDisburse.LoanId).Updates(disburse).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}

	// COMMIT THE CHANGE
	if err := tx.Commit().Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
