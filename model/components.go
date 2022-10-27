// Copyright (c) 2023 YaoBase
// 描述: 定义初始化相关保存配置参数的结构体
// 创建者: 张平
// 创建时间: 2022/5/26 17:12

package model

// LogData 日志参数相关结构体信息
type LogData struct {
	LogPath    string //日志文件目录
	LogLevel   string // 日志级别
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

// SslData 定义https相关配置信息
type SslData struct {
	HttpsEnable     bool   //是否开启https
	SslCrtPath      string //crt文件路径
	SslKeyPath      string //key文件路径
	SslProtocols    string //ssl支持的协议
	SslCipher       string //ssl支持的加密算法清单
	SslPasswordFile string
}

// PrometheusData Prometheus配置
type PrometheusData struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

// GrafanaData Grafana配置
type GrafanaData struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

// ConsulData Consul配置
type ConsulData struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}
