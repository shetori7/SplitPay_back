package usecase

import (
	"SplitPay_back/internal/domain"
	"SplitPay_back/internal/dto"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u *domain.Wari_user) {
	interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) GetInfo() []domain.Wari_user {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) GetInfoByGroupId(groupUuid string) []dto.UserByGroupIdDto {
	return interactor.UserRepository.SelectByGroupId(groupUuid)
}

func (interactor *UserInteractor) Delete(id int) {
	interactor.UserRepository.Delete(id)
}
