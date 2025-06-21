package domain

type Wari_user struct {
	UserID    int    `gorm:"primaryKey;autoIncrement" json:"user_id"`
	UserName  string `json:"user_name"`
	GroupUuid string `json:"group_uuid"`
}
