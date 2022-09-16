package dao

import (
	"fmt"
	"go-blog/models"
	"log"
)

// GetUserNameById 根据user id 获取username
func GetUserNameById(uid int) string {
	// 注意查询的时候只写一个字段即可
	row := DB.QueryRow("select user_name from blog_user where uid=?", uid)
	if row.Err() != nil {
		log.Println("根据id查询username出现错误: ", row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}

// GetUser 注意只有返回值为指针类型才可以return nil, 返回指针类型比较方便因为这个时候可以判断是否为nil判断用户是否存在
func GetUser(userName, password string) *models.User {
	fmt.Printf("查询的用户名与密码: ", userName, password)
	// limit 1保证查询的数据只有一条
	row := DB.QueryRow("select * from blog_user where user_name = ? and passwd = ? limit 1",
		userName, password)
	if row.Err() != nil {
		log.Println("查询用户信息错误: ", row.Err())
		return nil
	}
	// 注意user需要初始化才可以使用
	var user = &models.User{}
	// 将查询到的信息存储到user对应字段中, 注意需要一一对应
	err := row.Scan(&user.Uid, &user.UserName, &user.Password, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println("存储用户信息错误: ", err)
		return nil
	}
	fmt.Printf("user: %v", user)
	return user
}
