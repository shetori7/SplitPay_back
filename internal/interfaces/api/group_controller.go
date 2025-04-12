package controllers

import (
	"SplitPay_back/internal/domain"
	"SplitPay_back/internal/interfaces/database"
	"SplitPay_back/internal/usecase"
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

func (controller *GroupController) Create(groupname string) domain.Wari_group {
	g := domain.Wari_group{}
	g.GroupName = groupname

	controller.Interactor.Add(&g)
	return g
}
