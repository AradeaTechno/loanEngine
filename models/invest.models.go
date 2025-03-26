package models

import (
	"amarthaloan/config"
	"amarthaloan/db"
	"amarthaloan/helpers"
	"net/http"

	"gorm.io/gorm/clause"
)

type dataLoan struct {
	BorrowerId      int `json:"borrower_id"`
	PrincipalAmount int `json:"principal_amount"`
	Rate            int `json:"rate"`
	Roi             int `json:"roi"`
}

func GetOfferLoan() ([]dataLoan, int, error) {
	var dataLoans []dataLoan
	con := db.CreateConn()
	if err := con.Table("loans").
		Select("borrower_id, principal_amount, rate, roi").
		Where("state = ?", "approved").
		Find(&dataLoans).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return dataLoans, http.StatusOK, nil
}

func IsLoanOffer(loanId string) (string, int, int, error) {
	var objLoan helpers.LoanStruct
	con := db.CreateConn()
	if err := con.First(&objLoan, "loan_id = ? AND state = ?", loanId, "approved").Error; err != nil {
		return "", 0, http.StatusNotFound, err
	}
	return objLoan.PrincipalAmmount, objLoan.InvestedAmount, http.StatusOK, nil
}

func DoInvest(objLoan *helpers.LoanStruct, objInvest *helpers.InvestStruct, totalInvest int) (int, error) {
	con := db.CreateConn()
	tx := con.Begin()
	if tx.Error != nil {
		return http.StatusInternalServerError, tx.Error
	}

	// LOCK THE LOAN ROW
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("loan_id = ?", objInvest.LoanId).
		First(&objLoan).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}

	// CREATE DATA INVESTMENT
	if err := tx.Create(&objInvest).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}

	// UPDATE LOAN DATA
	if err := tx.Model(&objLoan).Where("loan_id = ?", objInvest.LoanId).Updates(objLoan).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}

	// IF STOCK PRINCIPAL == 0, CHANGE LOAN STATE TO INVESTED
	if totalInvest == 0 {
		appConf := config.AppConfig()
		agreementLink := appConf.APP_DOMAIN + ":" + appConf.APP_PORT + "/agreement-letter/" + objInvest.LoanId
		changeState := map[string]any{
			"state":            "invested",
			"agreement_letter": agreementLink,
		}
		if err := tx.Model(&objLoan).Where("loan_id = ?", objInvest.LoanId).Updates(changeState).Error; err != nil {
			tx.Rollback()
			return http.StatusInternalServerError, err
		}

		// SEND EMAIL
		var objInvestor helpers.InvestorStruct
		var emails []string
		if err := tx.Model(&objInvestor).
			Select("email").
			Joins("INNER JOIN loans_investment ON loans_investment.investor_id = investor.investor_id").
			Where("loans_investment.loan_id = ?", objInvest.LoanId).
			Pluck("email", &emails).Error; err != nil {
			tx.Rollback()
			return http.StatusInternalServerError, err
		}

		for _, email := range emails {
			emailData := new(helpers.EmailData)
			emailData.Email = email
			emailData.TypeEmail = "invested"
			emailData.AppName = appConf.APP_NAME
			emailData.Subject = "Investment Update"
			emailData.Link = agreementLink
			emailData.TeamEmail = "team@email.com"
			go helpers.SendEmail(*emailData)
		}
	}

	// COMMIT THE TRANSACTION
	if err := tx.Commit().Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
