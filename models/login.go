package models

type LoginResult struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"userInfo"`
}
