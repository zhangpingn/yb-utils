// Copyright (c) 2023 YaoBase
// 描述:	mysql连接对象初始化
// 创建者: 张平
// 创建时间: 2022/5/26 14:39

package mysql

import (
	"fmt"
	"github.com/zhangpingn/yb-utils/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var (
	// DB mysql对象实例
	DB *gorm.DB
)

// MysqlData Mysql相关结构体信息
type MysqlData struct {
	Host     string        `json:"host"`
	Port     uint          `json:"port"`
	Password string        `json:"password"`
	Username string        `json:"username"`
	DbName   string        `json:"db_name"`
	MaxIdle  int           `json:"max_idle"`
	MaxOpen  int           `json:"max_open"`
	ConnTTL  time.Duration `json:"conn_ttl"`
}

// InitMysql
// 描述：初始化mysql对象实例
// 创建者：张平
// 创建时间：2022/5/26 17:13
// 参数：
//	 [in]
//		无
//   [out]
//		无
//   [in/out]
//		无
func (mysqlData *MysqlData) InitMysql() {
	var err error

	for {
		// 拼装连接mysql的dsn，用于mysql连接传参
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			mysqlData.Username, mysqlData.Password,
			mysqlData.Host, mysqlData.Port, mysqlData.DbName)
		if DB, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,
			SkipInitializeWithVersion: true,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			PrepareStmt: true,
		}); err != nil {
			log.Error(err.Error())
			time.Sleep(time.Second * 2)
			continue
		}
		DB = DB.Debug()
		db, _ := DB.DB()
		// 判断数据库是否连接成功
		if err = db.Ping(); err != nil {
			// 需进行重连操作
			log.Error(err.Error())
			time.Sleep(time.Second * 2)
			continue
		}
		// 设置数据库连接池的最大空闲连接数量
		db.SetMaxIdleConns(mysqlData.MaxIdle) // default 2. If n <= 0, no idle connections are retained.
		// 设置数据库连接池的最大连接数量
		db.SetMaxOpenConns(mysqlData.MaxOpen) // default unlimited. If <= 0, there is no limit
		// 设置可重用连接的最长时间
		db.SetConnMaxLifetime(mysqlData.ConnTTL)
		break
	}

	log.Error("Mysql资源初始化完成")
}

// GetDB
// 描述：用于后续数据库操作时获取数据库操作实例对象
// 创建者：张平
// 创建时间：2022/5/26 17:13
// 参数：
//	 [in]
//		无
//   [out]
//		数据库链接对象实例
//   [in/out]
//		无
func (mysqlData *MysqlData) GetDB() *gorm.DB {
	db, _ := DB.DB()
	if err := db.Ping(); err != nil {
		// 需进行重连操作
		mysqlData.InitMysql()
	}
	return DB
}
