package domain

type Wari_group struct {
	GroupName string      `json:"group_name"`
	GroupUuid string      `json:"group_uuid"`
	Users     []Wari_user `gorm:"-"`
}
