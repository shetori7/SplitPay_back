package usecase

import (
	"SplitPay_back/internal/domain"

	"github.com/google/uuid"
)

type GroupInteractor struct {
	GroupRepository GroupRepository
}

func (interactor *GroupInteractor) Add(g *domain.Wari_group) {
	newUUID := uuid.New().String()
	g.GroupUuid = newUUID
	interactor.GroupRepository.Store(g)
}

func (interactor *GroupInteractor) GetInfo() []domain.Wari_group {
	return interactor.GroupRepository.Select()
}

func (interactor *GroupInteractor) Delete(id int) {
	interactor.GroupRepository.Delete(id)
}

func (interactor *GroupInteractor) GetInfoByGroupUuid(groupUuid string) domain.Wari_group {
	group := interactor.GroupRepository.SelectByGroupUuid(groupUuid)
	if group == nil {
		return domain.Wari_group{}
	}
	return *group
}
