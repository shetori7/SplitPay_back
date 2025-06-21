package database

import "SplitPay_back/internal/domain"

type GroupRepository struct {
	SqlHandler
}

func (db *GroupRepository) Store(g *domain.Wari_group) {
	db.Create(&g)
}

func (db *GroupRepository) Select() []domain.Wari_group {
	group := []domain.Wari_group{}
	db.FindAll(&group)
	return group
}

func (db *GroupRepository) Delete(id int) {
	group := []domain.Wari_user{}
	db.DeleteById(&group, id)
}

func (db *GroupRepository) SelectByGroupUuid(groupUuid string) *domain.Wari_group {
	group := &domain.Wari_group{}
	db.Raw().Table("wari_groups").Where("group_uuid = ?", groupUuid).First(&group)
	db.Raw().Table("wari_users").Where("group_uuid = ?", groupUuid).Find(&group.Users)
	db.Raw().Table("wari_final_payments").Where("group_uuid = ?", groupUuid).Find(&group.Payments)

	if group == nil {
		return nil
	}
	return group
}
