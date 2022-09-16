package service

import (
	"errors"
	"fmt"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/utils"
	"log"
)

func Login(userName, password string) (*models.LoginResult, error) {
	password = utils.Md5Crypt(password, "mszlu")
	fmt.Println("用户名与密码: ", userName, password)
	user := dao.GetUser(userName, password)
	if user == nil {
		return nil, errors.New("账号密码错误")
	}
	uid := user.Uid
	// 生成Token, 相当于是一个令牌(jwt技术), 前端拿到令牌之后那么过期了没有(临时的钥匙)
	// JWT: A.B.C
	token, err := utils.Award(&uid)
	if err != nil {
		log.Println("token未能够生成...")
	}
	var userInfo models.UserInfo
	userInfo.Uid = uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginResult{
		Token:    token,
		UserInfo: userInfo,
	}
	return lr, nil
}
