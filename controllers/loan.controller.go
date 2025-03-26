package controllers

import (
	"amarthaloan/helpers"
	"amarthaloan/models"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateLoan(c echo.Context) error {
	var objLoans helpers.LoanStruct

	objLoans.CreatedBy = c.FormValue("created_by")
	objLoans.BorrowerId = c.FormValue("borrower_id")
	objLoans.PrincipalAmmount = c.FormValue("principal_amount")
	objLoans.Rate = c.FormValue("rate")
	objLoans.Roi = c.FormValue("roi")
	objLoans.Sanitize()
	validate := helpers.ValidateStruct(objLoans)
	if len(validate) > 0 {
		for _, msg := range validate {
			return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New(msg), nil))
		}
	}

	// CHECK IS STAFF EXIST
	isStaffExists, err := models.IsStaffExists(objLoans.CreatedBy)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
	}
	if !isStaffExists {
		return c.JSON(http.StatusNotFound, helpers.ApiResponse(errors.New("staff not exists"), nil))
	}

	// CHECK IS BORROWER EXISES
	isBorrowerExists, err := models.IsBorrowerExists(objLoans.BorrowerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
	}
	if !isBorrowerExists {
		return c.JSON(http.StatusNotFound, helpers.ApiResponse(errors.New("borrower not exists"), nil))
	}

	// CREATE LOAN <> RETURN ERROR IF BORROWER HAS LOAN EXCEPT WITH DISBURSED STATE
	now := time.Now()
	objLoans.CreatedAt = &now
	objLoans.ModifiedAt = &now
	result, err := models.CreateLoan(&objLoans)
	if err != nil {
		return c.JSON(result, helpers.ApiResponse(err, nil))
	}
	return c.JSON(result, helpers.ApiResponse(errors.New("loan successfully created"), nil))
}

func GetLoanById(c echo.Context) error {
	loanId := c.Param("loanId")

	if loanId == "" {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("param loan can not be empty"), nil))
	}

	loans, code, err := models.GetLoanById(loanId, false)
	if err != nil {
		return c.JSON(code, helpers.ApiResponse(err, nil))
	}

	return c.JSON(code, helpers.ApiResponse(nil, loans))
}

func GetAllLoan(c echo.Context) error {
	res, code, err := models.GetAllLoan()
	if err != nil {
		return c.JSON(code, helpers.ApiResponse(err, nil))
	}
	return c.JSON(code, helpers.ApiResponse(nil, res))
}

func LoanProof(c echo.Context) error {
	var objProof helpers.ProofStruct

	loanId := c.FormValue("loan_id")
	if loanId == "" {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("loan_id is required"), nil))
	}
	isProposed, err := models.IsProposedExists(loanId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
	}
	if !isProposed {
		return c.JSON(http.StatusNotFound, helpers.ApiResponse(errors.New("loan not found"), nil))
	}

	intLoanId, err := strconv.Atoi(loanId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("invalid loan id type"), nil))
	}

	stafId := c.FormValue("staff_id")
	if stafId == "" {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("staff_id is required"), nil))
	}
	isStaffExists, err := models.IsStaffExists(stafId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
	}
	if !isStaffExists {
		return c.JSON(http.StatusNotFound, helpers.ApiResponse(errors.New("staff is not exists"), nil))
	}
	intStaffId, err := strconv.Atoi(stafId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("staff id invalid"), nil))
	}

	// HANDLING PICTURES
	uploadPath := "./uploads/proof"
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("error retrieving form"), nil))
	}

	files := form.File["proof_pictures"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("no files uploaded"), nil))
	}

	// CREATE DIRECTORY
	dirFile := filepath.Join(uploadPath, loanId)
	if _, err := os.Stat(dirFile); os.IsNotExist(err) {
		err = os.MkdirAll(dirFile, os.ModePerm)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
		}
	}

	// PROCESS FILE
	var errorList []string
	now := time.Now()
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
		}
		defer src.Close()

		// CHECK IS FILE ALLOWED
		err = helpers.IsAllowedImage(src)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.ApiResponse(err, nil))
		}

		// CREATE UNIQUE FILE NAME
		fileName := fmt.Sprintf("%s_%s", now.Format("20060102150405"), file.Filename)
		filePath := filepath.Join(dirFile, fileName)
		// CREATE DESTINATION INSIDE UPLOAD DIR
		dst, err := os.Create(filePath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
		}
		defer dst.Close()

		// COPY THE FILE CONTENT TO THE DESTINATION
		if _, err = io.Copy(dst, src); err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
		}

		objProof.LoanID = intLoanId
		objProof.StaffId = intStaffId
		objProof.ProofPic = filePath
		objProof.CreatedAt = &now

		err = models.AddProofPic(&objProof)
		if err != nil {
			customErrMsg := fmt.Sprintf("failed to upload image: %s):<br> Error : %v", file.Filename, err)
			errorList = append(errorList, customErrMsg)
			continue
		}
	}

	// IF ERROR OCCURED
	if len(errorList) > 0 {
		return c.JSON(http.StatusPartialContent, helpers.ApiResponse(errors.New("some file can not be proceed"), map[string]any{"error": errorList}))
	}

	return c.JSON(http.StatusOK, helpers.ApiResponse(errors.New("file successfully uploaded"), nil))

}

func GetLoanProof(c echo.Context) error {
	loanId := c.Param("loanId")
	if loanId == "" {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("loan_id is required"), nil))
	}
	detailLoan, code, err := models.GetLoanProof(loanId)
	if err != nil {
		return c.JSON(code, helpers.ApiResponse(err, nil))
	}
	return c.JSON(code, helpers.ApiResponse(nil, detailLoan))
}

func Approval(c echo.Context) error {
	// var objLoans helpers.LoanStruct

	loanId := c.FormValue("loan_id")
	if loanId == "" {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("loan_id is required"), nil))
	}

	// CHECK IS LOAN IN PROPOSE STATE
	isProposed, err := models.IsProposedExists(loanId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
	}
	if !isProposed {
		return c.JSON(http.StatusNotFound, helpers.ApiResponse(errors.New("loan not found"), nil))
	}
	intLoanId, err := strconv.Atoi(loanId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("staff id invalid"), nil))
	}

	// CHECK IS PROPOSED LOAN HAS PROOF
	isLoanHasProof, err := models.IsProofExists(loanId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
	}
	if !isLoanHasProof {
		return c.JSON(http.StatusForbidden, helpers.ApiResponse(errors.New("loan has no proof and can not to be approved"), nil))
	}

	// CHECK STAFF INPUT
	stafId := c.FormValue("staff_id")
	if stafId == "" {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("staff_id is required"), nil))
	}
	isStaffExists, err := models.IsStaffExists(stafId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
	}
	if !isStaffExists {
		return c.JSON(http.StatusNotFound, helpers.ApiResponse(errors.New("staff is not exists"), nil))
	}
	intStaffId, err := strconv.Atoi(stafId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("staff id invalid"), nil))
	}

	now := time.Now()
	dataUpdate := map[string]any{
		"state":         "approved",
		"approved_by":   intStaffId,
		"approval_date": now,
		"modified_at":   now,
	}

	updateLoan, err := models.UpdateLoan(intLoanId, dataUpdate)
	if err != nil {
		return c.JSON(updateLoan, helpers.ApiResponse(err, nil))
	}

	return c.JSON(updateLoan, helpers.ApiResponse(errors.New("loan successfully approved"), nil))

}

func Disburse(c echo.Context) error {
	var objDisburse helpers.DisburseStruct

	staffId := c.FormValue("staff_id")
	intStaffId, err := strconv.Atoi(staffId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("invalid staff_id format"), nil))
	}

	loanId := c.FormValue("loan_id")
	if loanId == "" {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("loan_id is required"), nil))
	}
	loan, code, err := models.GetLoanById(loanId, true)
	if err != nil {
		return c.JSON(code, helpers.ApiResponse(err, nil))
	}
	loanStruct, ok := loan.(helpers.LoanStruct)
	if !ok {
		return c.JSON(code, helpers.ApiResponse(errors.New("type assertion failed"), nil))
	}

	if loanStruct.State != "invested" {
		return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New("loan is not invested"), nil))
	}

	// SIGNED AGREEMENT HANDLE
	// CREATE DIRECTORY
	uploadPath := "./uploads/signed_agreement"
	dirFile := filepath.Join(uploadPath, loanId)
	if _, err := os.Stat(dirFile); os.IsNotExist(err) {
		err = os.MkdirAll(dirFile, os.ModePerm)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
		}
	}

	// GET FILE
	file, err := c.FormFile("signed_agreement")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(errors.New("invalid file"), nil))
	}
	// CHECK IS PDF
	if !helpers.IsAllowedFile("pdf", file) {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(errors.New("file should be in pdf format"), nil))
	}
	// OPENING FILE
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
	}
	defer src.Close()

	// CREATE UNIQUE FILE NAME
	fileName := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), file.Filename)
	filePath := filepath.Join(dirFile, fileName)
	// Create the destination file inside the user's directory
	dst, err := os.Create(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
	}
	defer dst.Close()

	// Copy the file content to the destination
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ApiResponse(err, nil))
	}
	// HANDLE FILE

	objDisburse.LoanId = loanStruct.ID
	objDisburse.StaffId = intStaffId
	objDisburse.SignAgreement = filePath
	objDisburse.Sanitize()
	validate := helpers.ValidateStruct(objDisburse)
	if len(validate) > 0 {
		for _, msg := range validate {
			return c.JSON(http.StatusBadRequest, helpers.ApiResponse(errors.New(msg), nil))
		}
	}

	now := time.Now()
	objDisburse.CreatedAt = &now

	disbured, err := models.Disburse(&objDisburse)
	if err != nil {
		return c.JSON(disbured, helpers.ApiResponse(err, nil))
	}
	return c.JSON(disbured, helpers.ApiResponse(errors.New("loan successfully disbursed"), nil))
}
