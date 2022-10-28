// Copyright (c) 2023 YaoBase
// 描述:	redis连接对象初始化
// 创建者: 张平
// 创建时间: 2022/5/26 14:39

package redis

import (
	"fmt"
	"github.com/zhangpingn/yb-utils/log"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

// 定义redis连接对象
var (
	redisClient *redis.Client
	wg          sync.WaitGroup
)

type RedisData struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

// InitRedis
// 描述：初始化Redis对象实例
// 创建者：张平
// 创建时间：2022/5/26 17:13
// 参数：
//	 [in]
//		无
//   [out]
//		无
//   [in/out]
//		无
func (redisData *RedisData) InitRedis() {
	wg.Add(1) // 这里原来的代码有错误，应该在主线程Add

	go func() {
		defer wg.Done()
		//wg.Add(1)
		//初始化redis
		redisClient = redisData.ConnRedis()
	}()
	wg.Wait()
}

// ConnRedis
// 描述：连接redis实际方法
// 创建者：张平
// 创建时间：2022/5/26 17:13
// 参数：
//	 [in]
//		ip：redis服务器ip
//		password：redis服务器密码
//		port：redis服务器端口
//   [out]
//		redis连接对象实例
//   [in/out]
//		无
func (redisData *RedisData) ConnRedis() *redis.Client {
	var (
		redisClient *redis.Client
		err         error
	)
	for {
		redisClient = redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%d", redisData.Host, redisData.Port),
			Password:     redisData.Password,
			DB:           0,
			PoolSize:     50,
			MinIdleConns: 10,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		})
		if _, err = (*redisClient).Ping().Result(); err != nil {
			log.Error(err.Error())
			time.Sleep(time.Second * 2)
			continue
		}
		log.Error(fmt.Sprintf("Redis【%s:%d】资源初始化完成", redisData.Host, redisData.Port))
		break
	}
	return redisClient
}

// GetClient
// 描述：获取连接，若连接中断，则进行重试连接操作，每2秒执行一次，直至重试连接成功
// 创建者：张平
// 创建时间：2022/5/26 17:13
// 参数：
//	 [in]
//		无
//   [out]
//		redis连接对象实例
//   [in/out]
//		无
func (redisData *RedisData) GetClient() *redis.Client {
	if _, err := redisClient.Ping().Result(); err != nil {
		redisClient = redisData.ConnRedis()
	}
	return redisClient
}
