package request

type GraoupNewRequestBody struct {
	GroupName string   `json:"group_name"`
	Users     []string `json:"users"`
}
