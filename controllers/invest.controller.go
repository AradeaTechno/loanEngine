package controllers

import (
	"amarthaloan/config"
	"amarthaloan/helpers"
	"amarthaloan/models"
	"bytes"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo/v4"
)

func OfferLoan(c echo.Context) error {
	offerLoan, code, err := models.GetOfferLoan()
	if err != nil {
		return c.JSON(code, helpers.ApiResponse(err, nil))
	}
	return c.JSON(code, offerLoan)
}

func DoInvest(c echo.Context) error {
	var objInvest helpers.InvestStruct
	var objLoan helpers.LoanStruct

	objInvest.InvestorId = c.FormValue("investor_id")
	objInvest.LoanId = c.FormValue("loan_id")
	objInvest.Amount = c.FormValue("amount")
	objInvest.Sanitize()
	validate := helpers.ValidateStruct(objInvest)
	if len(validate) > 0 {
		for _, msg := range validate {
			return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New(msg), nil))
		}
	}

	isInvestorExist, err := models.IsInvestorExists(objInvest.InvestorId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
	}
	if !isInvestorExist {
		return c.JSON(http.StatusNotFound, helpers.ApiResponse(errors.New("investor not exists"), nil))
	}

	// CHECK IS LOAN ALREADY APPROVED
	principalAmount, investedAmount, code, err := models.IsLoanOffer(objInvest.LoanId)
	if err != nil {
		return c.JSON(code, helpers.ApiResponse(err, nil))
	}

	intPrincipalAmount, err := strconv.Atoi(principalAmount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("invalid principal amount format"), nil))
	}

	intInvestedAmount, err := strconv.Atoi(objInvest.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("invalid amount format"), nil))
	}

	// TO CHECK LIMIT INVESTMENT
	stockPrincipal := intPrincipalAmount - investedAmount
	if intInvestedAmount > stockPrincipal {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("you can not invest more than : "+strconv.Itoa(stockPrincipal)), nil))
	}

	// CREATE FLAG TO CHECK IS INVESTED = PRINCIAP
	totalInvest := (investedAmount + intInvestedAmount) - intPrincipalAmount
	now := time.Now()
	finalInvestedAmount := intInvestedAmount + investedAmount
	objLoan.InvestedAmount = finalInvestedAmount
	objLoan.ModifiedAt = &now
	objInvest.CreatedAt = &now

	// PROCESS INVESTMENT
	code, err = models.DoInvest(&objLoan, &objInvest, totalInvest)
	if err != nil {
		return c.JSON(code, helpers.ApiResponse(err, nil))
	}

	return c.JSON(code, helpers.ApiResponse(errors.New("your investment succeed"), nil))
}

func GenerateAgreement(c echo.Context) error {
	loanId := c.Param("loanId")
	loan, code, err := models.GetLoanById(loanId, true)
	if err != nil {
		return c.JSON(code, helpers.ApiResponse(err, nil))
	}
	loanStruck, ok := loan.(helpers.LoanStruct)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(errors.New("type assertion failed"), nil))
	}

	appConf := config.AppConfig()
	// Create a new PDF instance
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a page
	pdf.AddPage()

	// Set font
	pdf.SetFont("Arial", "B", 16)

	// Add content
	pdf.Cell(40, 10, appConf.APP_NAME+" Invest Agreement For Loan "+loanStruck.PrincipalAmmount)

	pdf.SetFont("Arial", "", 12)
	pdf.Ln(10) // Move to the next line
	pdf.Cell(0, 10, "You have invested in our platform.")

	// Write the PDF to a buffer
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error generating PDF",
		})
	}

	// Set the response headers
	return c.Blob(http.StatusOK, "application/pdf", buf.Bytes())
}
