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
	controller.Interactor.Add(&u)
	createdUsers := controller.Interactor.GetInfo()
	c.JSON(201, createdUsers)
}

func (controller *UserController) CreateMultiple(userList []string, groupId int) []domain.Wari_user {
	users := make([]domain.Wari_user, len(userList))
	for i, userName := range userList {
		users[i] = domain.Wari_user{UserName: userName, GroupId: groupId}
	}

	for _, user := range users {
		controller.Interactor.Add(&user)
	}
	return users
}

func (controller *UserController) GetUser() []domain.Wari_user {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
