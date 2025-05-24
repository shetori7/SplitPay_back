package dto

type UserByGroupIdDto struct {
	GroupId   int    `json:"group_id"`
	GroupName string `json:"group_name"`
	GroupUuid string `json:"group_uuid"`
	UserId    int    `json:"user_id"`
	UserName  string `json:"user_name"`
}
