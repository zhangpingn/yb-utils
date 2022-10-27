// Copyright (c) 2023 YaoBase
// 描述: 定义md5加密
// 创建者: 张平
// 创建时间: 2022/5/26 17:13

package crypto

import (
	"crypto/md5"
	"fmt"
	"io"
)

// Md5Sum
// 描述：定义md5相关加密方法
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		raw：待加密的字符串
//   [out]
//		md5加密后的字符串
//   [in/out]
//		无
func Md5Sum(raw string) string {
	m := md5.New()
	io.WriteString(m, raw)
	return fmt.Sprintf("%x", m.Sum(nil))
}
