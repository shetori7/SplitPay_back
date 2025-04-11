package controllers

import (
	"SplitPay_back/internal/domain"
	"SplitPay_back/internal/interfaces/database"
	"SplitPay_back/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c *gin.Context) {
	u := domain.Wari_user{}
	c.Bind(&u)
	controller.Interactor.Add(u)
	createdUsers := controller.Interactor.GetInfo()
	c.JSON(201, createdUsers)
	return
}

func (controller *UserController) GetUser() []domain.Wari_user {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
