package domain

type Wari_group struct {
<<<<<<< HEAD
	GroupId   int         `gorm:"primaryKey;autoIncrement" json:"group_id"`
	GroupName string      `json:"group_name"`
	GroupUuid string      `json:"group_uuid"`
	Users     []Wari_user `gorm:"-"`
=======
	GroupName string `json:"group_name"`
	GroupUuid string `json:"group_uuid"`
>>>>>>> feature-20250621-conflict
}
