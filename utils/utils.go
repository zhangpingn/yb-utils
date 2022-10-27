// Copyright (c) 2023 YaoBase
// 描述: 封装一些公共的函数
// 创建者: 张平
// 创建时间: 2022/6/13 17:13

package utils

import (
	"YBFacadeService/common"
	"YBFacadeService/components/log"
	ybhttp "YBFacadeService/utils/http"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
)

// CreateSecret
// 描述：获取秘钥
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		n：获取密钥长度
//   [out]
//		密钥
//   [in/out]
//		无
func CreateSecret(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// StructToMap
// 描述：将结构体转换为map
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		obj：待转化的结构体对象实例
//   [out]
//		转化后的map对象实例
//   [in/out]
//		无
func StructToMap(obj interface{}) map[string]string {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]string)

	switch obj1.Kind() {
	case reflect.Ptr:
		for i := 0; i < obj1.Elem().NumField(); i++ {
			if obj2.Elem().Field(i).Interface() != nil {
				data[obj1.Elem().Field(i).Tag.Get("json")] = fmt.Sprintf("%v", obj2.Elem().Field(i).Interface())
			}
		}
	default:
		for i := 0; i < obj1.NumField(); i++ {
			if obj2.Field(i).Interface() != nil {
				data[obj1.Field(i).Tag.Get("json")] = fmt.Sprintf("%v", obj2.Field(i).Interface())
			}
		}
	}
	return data
}

// InArray
// 描述：判断一个字符串是否在一个切片中
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		key：待验证的字符串
//		arr：字符串切片
//   [out]
//		是否在切片中
//   [in/out]
//		无
func InArray(key string, arr []string) bool {
	for _, v := range arr {
		if key == v {
			return true
		}
	}
	return false
}

// InArrayInt
// 描述：判断一个int类型数值是否在一个int类型切片中
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		key：待验证的int类型数值
//		arr：int类型数值切片
//   [out]
//		是否在切片中
//   [in/out]
//		无
func InArrayInt(key int, arr []int) bool {
	for _, v := range arr {
		if key == v {
			return true
		}
	}

	return false
}

// InArrayUint16
// 描述：判断一个uint16类型数值是否在一个uint16类型切片中
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		key：待验证的uint16类型数值
//		arr：uint16类型数值切片
//   [out]
//		是否在切片中
//   [in/out]
//		无
func InArrayUint16(arr []uint16, key uint16) bool {
	for _, v := range arr {
		if v == key {
			return true
		}
	}
	return false
}

// Time2Second
// 描述：根据时间和单位转换成秒
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		time：时间数值
//		unit：单位
//		method：进行的操作类型
//   [out]
//		对应的秒数
//   [in/out]
//		无
func Time2Second(time int, unit, method string) int {
	var unitInt int
	switch unit {
	case "s":
		unitInt = 1
	case "m":
		unitInt = 60
	case "h":
		unitInt = 3600
	case "day":
		unitInt = 86400
	case "week":
		unitInt = 604800
	case "month":
		unitInt = 2592000
	case "year":
		unitInt = 31536000
	default:
		unitInt = 1
	}
	if method == "t2s" {
		return time * unitInt
	} else if method == "s2t" {
		return time / unitInt
	}
	return time
}

// Flow2Kb
// 描述：根据流量值和单位转换成相应的kb
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		flow：流量数值
//		unit：单位
//		method：进行的操作类型
//   [out]
//		对应的kb
//   [in/out]
//		无
func Flow2Kb(flow int64, unit, method string) int64 {
	var unitInt int64
	switch unit {
	case "b":
		unitInt = 1
	case "k":
		unitInt = 1024
	case "m":
		unitInt = 1048576
	case "g":
		unitInt = 1073741824
	case "t":
		unitInt = 1099511627776
	default:
		unitInt = 1
	}
	if method == "f2k" {
		return flow * unitInt
	} else if method == "k2f" {
		return flow / unitInt
	}
	return flow
}

// ParseIntToIpv4
// 描述：将整型数值转为ipv4字符串值
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		ip：整型ipv4数值
//   [out]
//		ip字符串值
//   [in/out]
//		无
func ParseIntToIpv4(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

// ParseIpv4toInt
// 描述：将整型数值转为ipv4字符串值
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		ip：字符串类型ipv4地址
//   [out]
//		整型ipv4地址
//   [in/out]
//		无
func ParseIpv4toInt(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

// ParseIpv6toInt
// 描述：字符串类型转整型ipv6地址
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		ip：字符串类型ipv6地址
//   [out]
//		整型ipv6地址
//   [in/out]
//		无
func ParseIpv6toInt(ip string) *big.Int {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To16())
	return ret
}

// ClientIP
// 描述：根据上下文获取请求客户端ip
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		r：http请求对象实例
//   [out]
//		请求客户端对象ip
//   [in/out]
//		无
func ClientIP(r *http.Request) string {

	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		str := r.RemoteAddr
		rs := []rune(str)
		ip = string(rs[:strings.LastIndex(str, ":")])
	}
	return ip
}

// GetMessageByErr
// 描述：根据错误返回提示信息
// 创建者：张平
// 创建时间：2022/6/15 13:34
// 参数：
//	 [in]
//		err：带解析的错误
//   [out]
//		提示信息
//   [in/out]
//		无
func GetMessageByErr(err error) string {
	if strings.HasPrefix(err.Error(), "Error 1062:") {
		return common.DataAlreadyExists
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return common.DataDoesNotExist
	} else if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		switch fieldErrors[0].Kind() {
		case reflect.String:
			return fieldErrors[0].Field() + " " + fieldErrors[0].Tag()
		case reflect.Int:
			return fieldErrors[0].Field() + " " + fieldErrors[0].Tag() + " " + fieldErrors[0].Param()
		}
		return fieldErrors[0].Field() + " " + fieldErrors[0].Tag()
	}
	return common.Failure
}

// CmdPing
// 描述：执行ping命令获取主机状态
// 创建者：张平
// 创建时间：2022/6/16 16:56
// 参数：
//	 [in]
//		host：主机ip
//   [out]
//		result：是否可以ping通
//		err：执行cmd命令过程中出现的错误
//   [in/out]
//		无
func CmdPing(host string) (result bool, err error) {
	sysType := runtime.GOOS
	if sysType == "linux" {
		cmd := exec.Command("/bin/sh", "-c", "ping -c 1 -w 10 "+host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		if strings.Contains(out.String(), "ttl=") {
			result = true
		}
	} else if sysType == "windows" {
		cmd := exec.Command("cmd", "/c", "ping -a -n 1 -w 10 "+host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()

		if strings.Contains(out.String(), "TTL=") {
			result = true
		}
	}
	return result, err
}

// GetCipherSuites
// 描述：定义用于tls1.2协议启动时配置的非对称加密算法相关
// 创建者：张平
// 创建时间：2022/5/26 17:13
// 参数：
//	 [in]
//		无
//   [out]
//		加密算法切片
//   [in/out]
//		无
func GetCipherSuites() []uint16 {
	var cipherSuites []uint16
	ciphers := strings.Split(common.Config.Ssl.SslCipher, ":")
	for _, v := range ciphers {
		switch v {
		case "ECDHE-ECDSA-AES256-GCM-SHA384":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384)

		/*case "kEDH+AESGCM":*/
		/*case "DHE-DSS-AES128-GCM-SHA256":*/
		/*case "DHE-RSA-AES128-GCM-SHA256":
		cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256)*/
		/*case "ECDHE-RSA-AES128-SHA256":
		cipherSuites = append(cipherSuites, tls.ECDHE_RSA_AES_128)*/
		/*case "ECDHE-ECDSA-AES128-SHA256":*/

		case "ECDHE-ECDSA-AES128-SHA":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA)
		case "ECDHE-RSA-AES128-SHA":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA)
		case "ECDHE-RSA-AES256-SHA384":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384)

		case "ECDHE-RSA-AES256-SHA":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA)
		case "ECDHE-ECDSA-AES256-SHA":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA)
			/*case "DHE-RSA-AES128-SHA256":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256)*/
			/*case "DHE-RSA-AES256-SHA256":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA)*/
		case "AES128-GCM-SHA256":
			cipherSuites = append(cipherSuites, tls.TLS_AES_128_GCM_SHA256)
		case "CHACHA20-POLY1305-SHA256":
			cipherSuites = append(cipherSuites, tls.TLS_CHACHA20_POLY1305_SHA256)
		case "AES256-GCM-SHA384":
			cipherSuites = append(cipherSuites, tls.TLS_AES_256_GCM_SHA384)
		case "ECDHE-RSA-CHACHA20-POLY1305":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305)
		case "ECDHE-ECDSA-CHACHA20-POLY1305":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305)
		case "ECDHE-RSA-AES256-GCM-SHA384":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384)
		case "ECDHE-ECDSA-AES256-SHA384":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384)
		case "ECDHE-RSA-AES128-CBC-SHA256":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256)
		case "ECDHE-RSA-AES128-CBC-SHA":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA)
		case "ECDHE-ECDSA-AES128-CBC-SHA256":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256)
		case "ECDHE-ECDSA-AES128-CBC-SHA":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA)
		case "ECDHE-RSA-AES256-CBC-SHA":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA)
		case "ECDHE-ECDSA-AES256-CBC-SHA":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA)
		case "RSA-AES256-GCM-SHA384":
			cipherSuites = append(cipherSuites, tls.TLS_RSA_WITH_AES_256_GCM_SHA384)
		case "RSA-AES128-GCM-SHA256":
			cipherSuites = append(cipherSuites, tls.TLS_RSA_WITH_AES_128_GCM_SHA256)
		case "ECDHE-ECDSA-AES128-GCM-SHA256":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256)
		case "ECDHE-RSA-AES128-GCM-SHA256":
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256)
		}
	}
	if cipherSuites == nil {
		cipherSuites = []uint16{
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		}
	} else {
		if !InArrayUint16(cipherSuites, tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256) {
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256)
		}
		if !InArrayUint16(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256) {
			cipherSuites = append(cipherSuites, tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256)
		}
	}
	return cipherSuites
}

// IsExist
// 描述：判断目录是否存在
// 创建者：张平
// 创建时间：2022/6/17 18:13
// 参数：
//	 [in]
//		dirName：目录名称
//   [out]
//		目录是否存在结果
//   [in/out]
//		无
func IsExist(dirName string) bool {
	_, err := os.Stat(dirName)
	return err == nil || os.IsExist(err)
}

// IsPicture
// 描述：判断是否是图片后缀
// 创建者：张平
// 创建时间：2022/6/20 10:34
// 参数：
//	 [in]
//		fileSuffix：文件后缀名
//   [out]
//		是否是图片后缀
//   [in/out]
//		无
func IsPicture(fileSuffix string) bool {
	return InArray(fileSuffix, []string{".jpg", ".bmp", ".jpeg", ".png", ".gif", ".pdf", ".webp"})
}

// ReloadRule
// 描述：更新告警规则配置
// 创建者：张平
// 创建时间：2022/7/22 10:06
// 参数：
//	 [in]
//		metricId：指标id
//		businessId：业务id
//   [out]
//		更新中出现的错误
//   [in/out]
//		无
func ReloadRule(metricId, businessId int) error {

	var paramMap = make(map[string]int)
	paramMap["business_id"] = businessId
	paramMap["metric_id"] = metricId
	reloadData, err := json.Marshal(paramMap)
	if err != nil {
		log.Error("序列化更新规则参数失败", zap.Error(err))
		return err
	}

	_, reloadRes, err := ybhttp.Put(fmt.Sprintf("%s/alarm/rule/reload", common.Config.MonitorHost), "", string(reloadData))
	if err != nil {
		log.Error(fmt.Sprintf("更新业务(%d)的指标类型为%d的规则失败", businessId, metricId), zap.Error(err))
		return err
	}

	// 返回deployRes的结果
	if reloadRes.Code != http.StatusOK {
		return errors.New(reloadRes.Message)
	}
	return nil
}

// ReloadAlertmanager
// 描述：更新alertmanager.yml
// 创建者：张平
// 创建时间：2022/7/22 10:07
// 参数：
//	 [in]
//		无
//   [out]
//		更新中出现的错误
//   [in/out]
//		无
func ReloadAlertmanager() error {

	_, reloadRes, err := ybhttp.Put(fmt.Sprintf("%s/alarm/alertmanger/reload", common.Config.MonitorHost), "", "")
	if err != nil {
		log.Error("更新alertmanager.yml失败", zap.Error(err))
		return err
	}

	// 返回deployRes的结果
	if reloadRes.Code != http.StatusOK {
		return errors.New(reloadRes.Message)
	}
	return nil
}

// GetDurationFormSecond
// 描述：根据秒获取相应格式的时间段字符串
// 创建者：张平
// 创建时间：2022/10/18 15:18
// 参数：
//	 [in]
//		无
//   [out]
//		更新中出现的错误
//   [in/out]
//		无
func GetDurationFormSecond(second int64) string {
	var durationKeys = []string{"week", "day", "h", "m", "s"}
	var durationValues = []int64{604800, 86400, 3600, 60, 1}
	var builder strings.Builder
	// 先求天
	var start = second
	var value int64
	for durationIndex, durationValue := range durationValues {
		value = start / durationValue
		start = start % durationValue
		if value == 0 {
			continue
		}
		builder.WriteString(fmt.Sprintf("%d%s", value, durationKeys[durationIndex]))
		if start == 0 {
			break
		}
	}

	return builder.String()
}
