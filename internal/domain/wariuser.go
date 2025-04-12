package domain

type Wari_user struct {
	UserID   int    `gorm:"primaryKey;autoIncrement" json:"user_id"`
	UserName string `json:"user_name"`
	GroupId  int    `json:"group_id"`
}
