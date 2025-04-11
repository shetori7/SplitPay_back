package controllers

import (
	"SplitPay_back/internal/domain"
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

// TODO：リクエストの構造体はここで持ってるでいいのか？
type RequestBody struct {
	GroupName string   `json:"group_name"`
	Users     []string `json:"users"`
}

func (controller *GroupController) Create(c *gin.Context) {
	var reqBody RequestBody

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	g := domain.Wari_group{
		GroupName: reqBody.GroupName,
	}
	controller.Interactor.Add(&g)
	c.JSON(http.StatusOK, gin.H{
		"groupUuid": g.GroupUuid,
		"groupName": g.GroupName,
		"groupId":   g.GroupId,
		"message":   "group created successfully",
	})
}
