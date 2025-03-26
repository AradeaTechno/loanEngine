package models

import (
	"amarthaloan/db"
	"amarthaloan/helpers"
	"errors"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func CreateLoan(objLoans *helpers.LoanStruct) (int, error) {
	con := db.CreateConn()

	borrowerId, err := strconv.ParseInt(objLoans.BorrowerId, 10, 64)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	createdBy, err := strconv.ParseInt(objLoans.CreatedBy, 10, 64)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if err := con.First(&objLoans, "borrower_id = ? AND state <> ?", objLoans.BorrowerId, "disbursed").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			dataLoan := map[string]any{
				"borrower_id":      borrowerId,
				"created_by_id":    createdBy,
				"principal_amount": objLoans.PrincipalAmmount,
				"rate":             objLoans.Rate,
				"roi":              objLoans.Roi,
				"state":            "proposed",
				"created_at":       objLoans.CreatedAt,
				"modified_at":      objLoans.ModifiedAt,
			}

			err := con.Model(&helpers.LoanStruct{}).Create(dataLoan).Error
			// err := con.Create(&objLoans).Error
			if err != nil {
				return http.StatusInternalServerError, err
			}
		} else {
			return http.StatusInternalServerError, err
		}
	} else {
		return http.StatusConflict, errors.New("borrower has active loan")
	}
	return http.StatusCreated, nil
}

func GetAllLoan() (any, int, error) {
	var objLoans helpers.LoanStruct
	var res helpers.Response

	con := db.CreateConn()
	if err := con.Find(&objLoans).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}
	return objLoans, http.StatusOK, nil
}

func GetLoanById(loanId string, complete bool) (any, int, error) {
	var objLoans helpers.LoanStruct
	con := db.CreateConn()
	if err := con.First(&objLoans, loanId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, err
		} else {
			return nil, http.StatusInternalServerError, err
		}
	}

	if !complete {
		dataLoan := map[string]any{
			"borrower_id":      objLoans.BorrowerId,
			"principal_amount": objLoans.PrincipalAmmount,
			"rate":             objLoans.Rate,
			"roi":              objLoans.Roi,
			"agreement_letter": objLoans.AgreementLetter,
		}
		return dataLoan, http.StatusOK, nil
	}

	return objLoans, http.StatusOK, nil
}

func IsProposedExists(loanId string) (bool, error) {
	var count int64
	var objLoans helpers.LoanStruct
	con := db.CreateConn()
	if err := con.First(&objLoans).Where("loan_id = ? AND state = ?", loanId, "proposed").Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// LOAN APPROVAL
func AddProofPic(objProof *helpers.ProofStruct) error {
	con := db.CreateConn()
	if err := con.Create(&objProof).Error; err != nil {
		return err
	}
	return nil
}

func IsProofExists(loanId string) (bool, error) {
	var count int64
	var objProof helpers.ProofStruct
	con := db.CreateConn()
	if err := con.First(&objProof).Where("loans_proof_id = ?", loanId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetLoanProof(loanId string) (map[string]any, int, error) {
	var objLoans helpers.LoanStruct
	var objProof []helpers.ProofStruct

	con := db.CreateConn()

	// GET LOAN
	if err := con.First(&objLoans, "loan_id = ?", loanId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, err
		} else {
			return nil, http.StatusInternalServerError, err
		}
	}

	// GET PROOF
	if err := con.Table("loans_proof").Where("loan_id = ?", loanId).Find(&objProof).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, err
		} else {
			return nil, http.StatusInternalServerError, err
		}
	}

	data := map[string]any{
		"loan":  objLoans,
		"proof": objProof,
	}

	return data, http.StatusOK, nil
}

func UpdateLoan(loanId int, dataUpdate map[string]any) (int, error) {
	var objLoans helpers.LoanStruct
	con := db.CreateConn()
	if err := con.Model(&objLoans).Where("loan_id = ?", loanId).Updates(dataUpdate).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
