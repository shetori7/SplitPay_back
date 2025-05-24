package controllers

import (
	"SplitPay_back/internal/domain"
	"SplitPay_back/internal/dto"
	"SplitPay_back/internal/infrastructure/request"
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

func (controller *UserController) CreateMultiple(c *gin.Context, groupId int) []domain.Wari_user {
	var reqBody request.GraoupNewRequestBody
	if value, exists := c.Get("request"); exists {
		reqBody = value.(request.GraoupNewRequestBody)
	}
	userList := reqBody.Users
	users := make([]domain.Wari_user, len(userList))
	for i, userName := range userList {
		users[i] = domain.Wari_user{UserName: userName, GroupId: groupId}
	}

	for i := range users {
		controller.Interactor.Add(&users[i])
	}
	return users
}

func (controller *UserController) GetUser() []domain.Wari_user {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) GetUserByGroupId(c *gin.Context) []dto.UserByGroupIdDto {
	var reqBody request.GetUserByGroupIdRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return nil
	}
	groupUuid := reqBody.GroupUuid
	c.Set("request", reqBody)
	res := controller.Interactor.GetInfoByGroupId(groupUuid)
	return res
}

func (controller *UserController) Delete(id int) {
	controller.Interactor.Delete(id)
}
