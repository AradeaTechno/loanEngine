package routes

import (
	"amarthaloan/config"
	"amarthaloan/controllers"
	"amarthaloan/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	appConf := config.AppConfig()
	api := appConf.API_PATTERN + "/" + appConf.API_VERSION
	e := echo.New()

	// TESTER
	e.GET(api+"/", func(c echo.Context) error {
		return c.String(http.StatusOK, "AmarthaLoan is up!")
	}, middleware.RateLimitMiddleware)

	staffGroup := e.Group(api+"/staff", middleware.RateLimitMiddleware)
	staffGroup.POST("/create", controllers.CreateStaff)

	investorGroup := e.Group(api+"/investor", middleware.RateLimitMiddleware)
	investorGroup.POST("/create", controllers.CreateInvestor)

	borrowerGroup := e.Group(api+"/borrower", middleware.RateLimitMiddleware)
	borrowerGroup.POST("/create", controllers.CreateBorrower)

	loanGroup := e.Group(api+"/loan", middleware.RateLimitMiddleware)
	loanGroup.POST("/create", controllers.CreateLoan)
	loanGroup.POST("/proof", controllers.LoanProof)
	loanGroup.PATCH("/approval", controllers.Approval)
	loanGroup.POST("/disburse", controllers.Disburse)
	loanGroup.GET("/proof/:loanId", controllers.GetLoanProof)
	loanGroup.GET("/list", controllers.GetAllLoan)
	loanGroup.GET("/list/:loanId", controllers.GetLoanById)

	investGroup := e.Group(api+"/investment", middleware.RateLimitMiddleware)
	investGroup.GET("/offered", controllers.OfferLoan)
	investGroup.POST("/do-invest", controllers.DoInvest)

	e.GET("/agreement-letter/:loanId", controllers.GenerateAgreement)

	return e
}
