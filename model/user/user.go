// Copyright (c) 2023 YaoBase
// 描述:	用户信息结构体以及数据库校验相关逻辑定义
// 创建者: 张平
// 创建时间: 2022/6/6 13:18

package user

import (
	"github.com/zhangpingn/yb-utils/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	UserName     string `json:"user_name" gorm:"primaryKey"`
	Password     string `json:"password"`
	UserFullName string `json:"user_full_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
}

type UserLoginRes struct {
	User
	PrimaryBusiness int `json:"primary_business"`
	Version         string
}

// Check
// 描述：数据库校验用户名密码是否正确
// 创建者：张平
// 创建时间： 2022/6/6 13:25
// 参数：
//	 [in]
//		无
//   [out]
//		correct：验证结果是否正确
//   [in/out]
//		this.Username：待验证的用户名
//		this.Password：待验证的密码
//		this.UserFullName：用户的姓名，用于前端页面展示
func (user *User) Check(db *gorm.DB) (correct bool) {

	password := strings.ToUpper(user.Password)
	err := db.Where("user_name=?", user.UserName).First(&user).Error
	if err != nil {
		log.Error("根据用户名查询用户信息失败, 失败原因：", zap.Error(err))
		return
	}

	// 判断用户名和密码与待验证的是否一致
	correct = password == user.Password

	return
}
