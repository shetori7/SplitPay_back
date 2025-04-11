package usecase

import (
	"SplitPay_back/internal/domain"

	"github.com/google/uuid"
)

type GroupInteractor struct {
	GroupRepository GroupRepository
}

func (interactor *GroupInteractor) Add(g domain.Wari_group) string {
	newUUID := uuid.New().String()
	g.GroupUuid = newUUID
	interactor.GroupRepository.Store(g)
	return newUUID
}

func (interactor *GroupInteractor) GetInfo() []domain.Wari_group {
	return interactor.GroupRepository.Select()
}

func (interactor *GroupInteractor) Delete(id string) {
	interactor.GroupRepository.Delete(id)
}
