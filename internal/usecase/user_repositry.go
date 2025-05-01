package usecase

import (
	"SplitPay_back/internal/domain"
)

type UserRepository interface {
	Store(*domain.Wari_user)
	Select() []domain.Wari_user
	Delete(id int)
}
