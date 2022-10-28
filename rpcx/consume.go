// Copyright (c) 2023 YaoBase
// 描述: 定义消费者方法
// 创建者: 张平
// 创建时间: 2022/10/13 13:10

package rpcx

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"github.com/zhangpingn/yb-utils/log"
)

// Consume
// 描述：调用服务提供者的方法进行消费
// 创建者：张平
// 创建时间：2022/10/13 13:23
// 参数：
//	 [in]
//		basePath：服务前缀路径
//		servicePath：服务路径
//		serviceMethod：服务方法
//		param：请求参数对象
//		res：返回结果对象
//   [out]
//		无
//   [in/out]
//		无
func Consume(basePath, servicePath, serviceMethod, consulAddr string, param, res interface{}) error {
	d := client.NewConsulDiscovery(basePath, servicePath, []string{consulAddr}, nil)
	xclient := client.NewXClient(servicePath, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer func() {
		if err := xclient.Close(); err != nil {
			log.Error("关闭xclient对象出错")
		}
	}()

	// xclient.Auth()进行请求token校验

	return xclient.Call(context.Background(), serviceMethod, &param, &res)
}
