package models

import (
	"amarthaloan/config"
	"amarthaloan/db"
	"amarthaloan/helpers"
	"errors"
	"net/http"
)

func GetOfferLoan() (helpers.Response, int, error) {
	var objLoan []helpers.LoanStruct
	var res helpers.Response
	con := db.CreateConn()
	if err := con.Where("state = ?", "approved").Find(&objLoan).Error; err != nil {
		return res, http.StatusInternalServerError, err
	}
	return helpers.ApiResponse(errors.New("data retrieved"), objLoan), http.StatusOK, nil
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
		// TAKE ALL INVESTORS
		var objInvestor helpers.InvestorStruct
		var emails []string
		if err := tx.Model(&objInvestor).
			Select("email").
			Joins("INNER JOIN loans_investment ON loans_investment.investor_id = investor.investor_id").
			Where("loans_investment.loan_id = ? ", objInvest.LoanId).
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

	// COMMIT THE CHANGE
	if err := tx.Commit().Error; err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
