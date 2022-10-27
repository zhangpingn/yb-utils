// Copyright (c) 2023 YaoBase
// 描述: 定义base64加密
// 创建者: 张平
// 创建时间: 2022/5/26 17:13

package crypto

import (
	"encoding/base64"
	"reflect"
	"unsafe"
)

// Encode
// 描述: 加密函数
// 创建者:曹壮
// 创建时间: 16:55 2022/6/30
// 参数:
//	 [in]
//		data: 待加密字符串
//   [out]
//		无
//   [in/out]
//		无
func Encode(data string) string {
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	coder := base64.NewEncoding(getBase64Table())
	return coder.EncodeToString(content)
}

// Decode
// 描述: 解密函数
// 创建者:曹壮
// 创建时间: 16:56 2022/6/30
// 参数:
//	 [in]
//		data: 已加密字符串
//   [out]
//		无
//   [in/out]
//		无
func Decode(data string) string {
	coder := base64.NewEncoding(getBase64Table())
	result, _ := coder.DecodeString(data)
	return *(*string)(unsafe.Pointer(&result))
}

// getBase64Table
// 描述: 获取Base64表
// 创建者:曹壮
// 创建时间: 16:57 2022/6/30
// 参数:
//	 [in]
//		无
//   [out]
//		无
//   [in/out]
//		无
func getBase64Table() string {
	str := "LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA"
	return str
}
