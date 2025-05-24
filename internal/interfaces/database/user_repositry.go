package database

import (
	"SplitPay_back/internal/domain"
	"SplitPay_back/internal/dto"
)

type UserRepository struct {
	SqlHandler
}

func (db *UserRepository) Store(u *domain.Wari_user) {
	db.Create(u)
}

func (db *UserRepository) Select() []domain.Wari_user {
	user := []domain.Wari_user{}
	db.FindAll(&user)
	return user
}

func (db *UserRepository) SelectByGroupId(groupUuid string) []dto.UserByGroupIdDto {
	userByGroupIdDto := []dto.UserByGroupIdDto{}
	db.Raw().Table("wari_users").
		Select("wari_users.*", "wari_groups.*").
		Joins("JOIN wari_groups ON wari_users.group_id = wari_groups.group_id").
		Where("wari_users.group_uuid = ?", groupUuid).
		Scan(&userByGroupIdDto)

	return userByGroupIdDto
}

func (db *UserRepository) Delete(id int) {
	user := []domain.Wari_user{}
	db.DeleteById(&user, id)
}
