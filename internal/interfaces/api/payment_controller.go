package controllers

import (
	"SplitPay_back/internal/infrastructure/request"
	"SplitPay_back/internal/interfaces/database"
	"SplitPay_back/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
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
	var reqBody request.PaymentNewRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := controller.Interactor.Add(reqBody.GroupId, reqBody.PayerId, reqBody.Amount, reqBody.ParticipantIds)
	controller.Interactor.ReCalcFinalPayment(reqBody.GroupId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Payment created successfully",
		})
		return
	}
}
