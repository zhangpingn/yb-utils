// Copyright (c) 2023 YaoBase
// 描述: 定义sha1加密
// 创建者: 张平
// 创建时间: 2022/5/26 17:13

package crypto

import (
	"crypto/sha1"
	"fmt"
	"io"
)

// Sha1Sum
// 描述：定义sha1加密
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		raw：待加密的字符串
//   [out]
//		sha1加密后的字符串
//   [in/out]
//		无
func Sha1Sum(raw string) string {
	s := sha1.New()
	io.WriteString(s, raw)
	return fmt.Sprintf("%x", s.Sum(nil))
}
