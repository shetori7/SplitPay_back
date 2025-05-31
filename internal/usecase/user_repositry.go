package usecase

import (
	"SplitPay_back/internal/domain"
	"SplitPay_back/internal/dto"
)

type UserRepository interface {
	Store(*domain.Wari_user)
	Select() []domain.Wari_user
	Delete(id int)
	SelectByGroupId(groupUuid string) []dto.UserByGroupIdDto
}
