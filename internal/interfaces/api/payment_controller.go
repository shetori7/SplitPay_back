package controllers

import (
	"SplitPay_back/internal/infrastructure/request"
	"SplitPay_back/internal/interfaces/database"
	"SplitPay_back/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
)

type PaymentContlloer struct {
	Interactor usecase.PaymentInteractor
}

func NewPaymentController(sqlHandler database.SqlHandler) *PaymentContlloer {
	return &PaymentContlloer{
		Interactor: usecase.PaymentInteractor{
			WariPaymentRepository: &database.WariPaymentRepository{
				SqlHandler: sqlHandler,
			},
			WariLoanRepository: &database.WariLoanRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *PaymentContlloer) Create(c *gin.Context) {
	var errors error
	var reqBody request.PaymentNewRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	addErr := controller.Interactor.Add(reqBody.GroupUuid, reqBody.PayerId, reqBody.Amount, reqBody.ParticipantIds)
	errors = multierror.Append(errors, addErr)
	reCalcErr := controller.Interactor.ReCalcFinalPayment(reqBody.GroupUuid)
	errors = multierror.Append(errors, reCalcErr)

	if finalErr := errors.(*multierror.Error).ErrorOrNil(); finalErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": finalErr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Payment created successfully",
	})
	return
}
