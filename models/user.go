package models

import "time"

type User struct {
	Uid      int       `json:"uid"`
	UserName string    `json:"userName"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

// UserInfo 用户信息
type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
