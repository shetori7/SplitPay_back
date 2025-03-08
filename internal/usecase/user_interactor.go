package usecase

import "SplitPay_back/internal/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.Wari_user) {
	interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) GetInfo() []domain.Wari_user {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) Delete(id string) {
	interactor.UserRepository.Delete(id)
}
