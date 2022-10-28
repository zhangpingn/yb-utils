// Copyright (c) 2023 YaoBase
// 描述:	获取yaobase数据库连接
// 创建者: 张平
// 创建时间: 2022/6/15 16:56

package yaobase

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/patrickmn/go-cache"
	"github.com/zhangpingn/yb-utils/crypto"
	"github.com/zhangpingn/yb-utils/log"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
	"time"
)

var (
	ybCache      = cache.New(5*time.Minute, 10*time.Minute)
	ybBusiness   = "yb:business:%s:%d"
	maxIdle      = 10
	maxOpen      = 20
	ybExpireTime = time.Minute * 20
)

type business struct {
	Id         int    `json:"id" gorm:"primaryKey;column:id"`
	YBHost     string `json:"yb_host" binding:"required"`
	YBPort     int    `json:"yb_port" binding:"required,min=1"`
	YBUser     string `json:"yb_user" binding:"required"`
	YBPassword string `json:"yb_password" binding:"required"`
	Status     int    `json:"status"`
	OtherSsIps string `json:"-"`
}

// GetDBByBusiness
// 描述：根据业务id获取yaobase数据库连接对象实例
// 创建者：张平
// 创建时间：2022/5/26 17:13
// 参数：
//	 [in]
//		businessId：业务id
//   [out]
//		db：yaobase数据库连接对象实例
//   [in/out]
//		无
func GetDBByBusiness(db *gorm.DB, businessId int, flags ...bool) (ybDB *sql.DB) {
	var (
		err error
	)

	if businessId == 0 {
		return nil
	}

	business := business{
		Id: businessId,
	}

	if err = db.First(&business).Error; err != nil {
		log.Error("根据业务id获取业务实例对象失败，失败原因：", zap.Error(err))
		return nil
	}

	// 该逻辑给部署使用
	if len(flags) <= 0 {
		if business.Status != 1 {
			log.Error("该业务未部署成功")
			return nil
		}
	}

	if ybDBVal, ok := ybCache.Get(fmt.Sprintf(ybBusiness, business.YBHost, business.YBPort)); ok && ybDBVal != nil {
		ybDB = ybDBVal.(*sql.DB)
		if err = ybDB.Ping(); err == nil {
			return ybDB
		}
		ybCache.Delete(fmt.Sprintf(ybBusiness, business.YBHost, business.YBPort))
	}

	var otherSsIps = []string{business.YBHost}
	if business.OtherSsIps != "" {
		otherSsIps = append(otherSsIps, strings.Split(business.OtherSsIps, ",")...)
	}
	password := crypto.Decode(business.YBPassword)
	var flag bool
	for _, v := range otherSsIps {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/TANG?charset=utf8mb4&parseTime=True",
			business.YBUser, password, v, business.YBPort)

		ybDB, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Error("根据业务id获取yaobase数据库实例对象失败，失败原因：", zap.Error(err))
			continue
		}
		//与数据库连接
		err = ybDB.Ping()
		if err != nil {
			log.Error("根据业务id获取yaobase数据库实例对象失败，失败原因：", zap.Error(err))
			continue
		}
		ybDB.SetMaxIdleConns(maxIdle)
		ybDB.SetMaxOpenConns(maxOpen)
		ybDB.SetConnMaxLifetime(ybExpireTime)
		flag = true
		break
	}

	if !flag {
		return nil
	}

	defer ybCache.Set(fmt.Sprintf(ybBusiness, business.YBHost, business.YBPort), ybDB, ybExpireTime)
	// 查询所有的sqlserver服务的节点ip，更新business表
	rows, _ := ybDB.Query("select svr_ip from __all_server where svr_type='yaosqlsvr'")
	otherSsIps = make([]string, 0)
	for rows.Next() {
		var svrIp string
		if err := rows.Scan(&svrIp); err != nil || svrIp == business.YBHost {
			continue
		}
		otherSsIps = append(otherSsIps, svrIp)
	}
	otherSsIpStr := strings.Join(otherSsIps, ",")
	business.OtherSsIps = otherSsIpStr
	db.Updates(business)
	return

}
