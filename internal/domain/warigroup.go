package domain

type Wari_group struct {
	GroupId   int         `gorm:"primaryKey;autoIncrement" json:"group_id"`
	GroupName string      `json:"group_name"`
	GroupUuid string      `json:"group_uuid"`
	Users     []Wari_user `gorm:"-"`
}
