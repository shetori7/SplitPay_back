package usecase

import (
	"SplitPay_back/internal/domain"
)

type GroupRepository interface {
	Store(*domain.Wari_group)
	Select() []domain.Wari_group
	Delete(id int)
	SelectByGroupUuid(groupUuid string) *domain.Wari_group
}
