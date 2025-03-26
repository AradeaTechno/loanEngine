package helpers

import (
	"time"

	"github.com/microcosm-cc/bluemonday"
)

var p = bluemonday.StrictPolicy()

// =============== STAFF =========================== //
type StaffStruct struct {
	StaffId    int        `json:"staff_id" gorm:"column:staff_id;primaryKey"`
	Name       string     `json:"staff_name" gorm:"column:name;not null" validate:"required"`
	Email      string     `json:"staff_email" gorm:"column:email;unique;not null" validate:"required,email"`
	Role       string     `json:"staff_role" gorm:"column:role;not null" validate:"required,role_validator"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at;not null"`
	ModifiedAt *time.Time `json:"modified_at" gorm:"column:modified_at;not null"`
}

func (StaffStruct) TableName() string {
	return "staff"
}

func (s *StaffStruct) Sanitize() {
	s.Name = p.Sanitize(s.Name)
	s.Email = p.Sanitize(s.Email)
	s.Role = p.Sanitize(s.Role)
}

// =================== BORROWER ========================= //
type BorrowerStruct struct {
	BorrowerId int        `json:"borrower_id" gorm:"column:borrower_id;primaryKey"`
	Name       string     `json:"borrower_name" gorm:"column:name" validate:"required"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at"`
	ModifiedAt *time.Time `json:"modified_at" gorm:"column:modified_at"`
}

func (BorrowerStruct) TableName() string {
	return "borrower"
}

func (s *BorrowerStruct) Sanitize() {
	s.Name = p.Sanitize(s.Name)
}

// ===================== INVESTOR ====================== //
type InvestorStruct struct {
	InvestorId int        `json:"investor_id" gorm:"column:investor_id;primaryKey"`
	Name       string     `json:"investor_name" gorm:"column:name;nol null" validate:"required"`
	Email      string     `json:"investor_email" gorm:"column:email;unique;nol null" validate:"required,email"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at;nol null"`
	ModifiedAt *time.Time `json:"modified_at" gorm:"column:modified_at;nol null"`
}

func (InvestorStruct) TableName() string {
	return "investor"
}

func (s *InvestorStruct) Sanitize() {
	s.Name = p.Sanitize(s.Name)
	s.Email = p.Sanitize(s.Email)
}

// ====================== LOAN ======================== //
type LoanStruct struct {
	ID               int        `json:"loan_id" gorm:"column:loan_id;primaryKey"`
	CreatedBy        string     `json:"created_by" gorm:"column:created_by_id;not null" validate:"required,numeric"`
	BorrowerId       string     `json:"borrower_id" gorm:"column:borrower_id;not null" validate:"required,numeric"`
	PrincipalAmmount string     `json:"principal_amount" gorm:"column:principal_amount;not null" validate:"required,number"`
	Rate             string     `json:"rate" gorm:"column:rate;not null" validate:"required,number"`
	Roi              string     `json:"roi" gorm:"column:roi;not null" validate:"required,number"`
	State            string     `json:"state" gorm:"column:state;not null;default:proposed"`
	AgreementLetter  string     `json:"agreement_letter" gorm:"column:agreement_letter"`
	ApprovedBy       string     `json:"approved_by" gorm:"column:approved_by"`
	ApprovalDate     *time.Time `json:"approval_time" gorm:"column:approval_date"`
	InvestedAmount   int        `json:"invested_amount" gorm:"column:invested_amount"`
	DisbursedBy      string     `json:"disbursed_by" gorm:"column:disbursed_by"`
	DisbursementDate *time.Time `json:"disbursement_date" gorm:"column:disbursement_date"`
	CreatedAt        *time.Time `json:"created_at" gorm:"column:created_at;not null"`
	ModifiedAt       *time.Time `json:"modified_at" gorm:"modified_at;not null"`
}

func (LoanStruct) TableName() string {
	return "loans"
}

func (s *LoanStruct) Sanitize() {
	s.CreatedBy = p.Sanitize(s.CreatedBy)
	s.BorrowerId = p.Sanitize(s.BorrowerId)
	s.PrincipalAmmount = p.Sanitize(s.PrincipalAmmount)
	s.Rate = p.Sanitize(s.Rate)
	s.Roi = p.Sanitize(s.Roi)
	s.AgreementLetter = p.Sanitize(s.AgreementLetter)
}

type ProofStruct struct {
	// ID           int        `json:"id" gorm:"column:loans_approval_id;;primaryKey;autoIncrement"`
	LoanID    int        `json:"loan_id" gorm:"column:loan_id;not null"`
	StaffId   int        `json:"staff_id" gorm:"column:staff_id;not null"`
	ProofPic  string     `json:"proof_picture" gorm:"column:proof_picture"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}

func (ProofStruct) TableName() string {
	return "loans_proof"
}

func (s *ProofStruct) Sanitize() {
	s.ProofPic = p.Sanitize(s.ProofPic)
}

// ================== INVEST ========================= //
type InvestStruct struct {
	ID         int        `json:"invest_id" gorm:"column:loan_investment_id;primaryKey;autoIncrement"`
	InvestorId string     `json:"investor_id" gorm:"column:investor_id" validate:"required,numeric"`
	LoanId     string     `json:"loan_id" gorm:"column:loan_id" validate:"required,numeric"`
	Amount     string     `json:"amount" gorm:"column:invest_amount" validate:"required,numeric"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at"`
}

func (InvestStruct) TableName() string {
	return "loans_investment"
}

func (s *InvestStruct) Sanitize() {
	s.LoanId = p.Sanitize(s.LoanId)
	s.InvestorId = p.Sanitize(s.InvestorId)
	s.Amount = p.Sanitize(s.Amount)
}

// ================== DISBURSE =========================== //
type DisburseStruct struct {
	ID            int        `json:"disburse_id" gorm:"column:loans_disburse_id;primaryKey"`
	LoanId        int        `json:"loan_id" gorm:"column:loan_id;not null"`
	StaffId       int        `json:"staff_id" gorm:"column:staff_id;not null"`
	SignAgreement string     `json:"sign_agreement" gorm:"column:signed_agreement;not null"`
	CreatedAt     *time.Time `json:"created_at" gorm:"column:created_at;not null"`
}

func (DisburseStruct) TableName() string {
	return "loans_disburse"
}

func (s *DisburseStruct) Sanitize() {
	s.SignAgreement = p.Sanitize(s.SignAgreement)
}
