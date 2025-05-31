package controllers

import (
	"SplitPay_back/internal/domain"
	"SplitPay_back/internal/infrastructure/request"
	"SplitPay_back/internal/interfaces/database"
	"SplitPay_back/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	Interactor usecase.GroupInteractor
}

func NewGroupController(sqlHandler database.SqlHandler) *GroupController {
	return &GroupController{
		Interactor: usecase.GroupInteractor{
			GroupRepository: &database.GroupRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *GroupController) Create(c *gin.Context) domain.Wari_group {
	g := domain.Wari_group{}
	var reqBody request.GraoupNewRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return g
	}
	c.Set("request", reqBody)
	g.GroupName = reqBody.GroupName
	controller.Interactor.Add(&g)
	return g
}

func (controller *GroupController) GetGroupByUuId(c *gin.Context) domain.Wari_group {
	groupUuId := c.Query("groupUuId")
	group := controller.Interactor.GetInfoByGroupUuid(groupUuId)
	return group
}
