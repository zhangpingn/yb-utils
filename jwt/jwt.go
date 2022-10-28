// Copyright (c) 2023 YaoBase
// 描述:	声明jwt相关变量以及创建token、验证token相关方法
// 创建者: 张平
// 创建时间: 2022/5/26 14:39

package jwt

import (
	"github.com/zhangpingn/yb-utils/model/user"
	"github.com/zhangpingn/yb-utils/rpcx"
)

type TokenParam struct {
	Username string `json:"username"`
	Ip       string `json:"ip"`
	Token    string `json:"token"`
}

// CreateToken
// 描述：创建token方法
// 创建者：张平
// 创建时间：2022/5/26 17:13
// 参数：
//	 [in]
//		targetUser：当前请求token的用户名
//		ip：请求创建的客户端ip
//   [out]
//		jwt：创建的token信息
//		err：创建token过程中出现的错误
//   [in/out]
//		无
func CreateToken(user *user.User, basePath, servicePath, serviceMethod, consulAddr string) (token *string, err error) {
	err = rpcx.Consume(basePath, servicePath, serviceMethod, consulAddr, user, token)
	return
}

// ValidateJWT
// 描述：验证jwt
// 创建者：张平
// 创建时间：2022/5/26 17:13
// 参数：
//	 [in]
//		jwt：待验证的token
//		ip：当前请求验证的客户端ip
//   [out]
//		验证的结果对象实例
//   [in/out]
//		无
func ValidateJWT(tokenParam *TokenParam, basePath, servicePath, serviceMethod, consulAddr string) (token *string, err error) {
	err = rpcx.Consume(basePath, servicePath, serviceMethod, consulAddr, tokenParam, token)
	return
}

// GetUserInfoByToken
// 描述：根据上下文获取当前操作用户信息
// 创建者：张平
// 创建时间：2022/5/26 17:13
// 参数：
//	 [in]
//		ctx：请求的上下文对象
//   [out]
//		结果1：当前操作用户对象实例
//		结果2：根据上下文获取当前操作用户信息过程中出现的错误
//   [in/out]
//		无
func GetUserInfoByToken(token *string, basePath, servicePath, serviceMethod, consulAddr string) (user *user.User, err error) {
	err = rpcx.Consume(basePath, servicePath, serviceMethod, consulAddr, token, user)
	return
}
