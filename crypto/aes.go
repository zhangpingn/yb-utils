// Copyright (c) 2023 YaoBase
// 描述: 定义aes相关加密解密方法
// 创建者: 张平
// 创建时间: 2022/5/26 17:13

package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/zhangpingn/yb-utils/log"
)

// paddingText1
// 描述：填充字符串
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		str：待填充的byte切片
//		blockSize：block大小
//   [out]
//		填充后的byte切片
//   [in/out]
//		无
func paddingText1(str []byte, blockSize int) []byte {
	//需要填充的数据长度
	paddingCount := blockSize - len(str)%blockSize
	//填充数据为：paddingCount ,填充的值为：paddingCount
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newPaddingStr := append(str, paddingStr...)
	//fmt.Println(newPaddingStr)
	return newPaddingStr
}

// unPaddingText1
// 描述：去掉填充字符
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		str：待去填充byte切片
//   [out]
//		去掉填充字符后的byte切片
//   [in/out]
//		无
func unPaddingText1(str []byte) []byte {
	n := len(str)
	count := int(str[n-1])
	newPaddingText := str[:n-count]
	return newPaddingText
}

// EncryptAES
// 描述：加密
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		src：待加密byte切片
//		key：密钥
//   [out]
//		加密后的byte切片
//   [in/out]
//		无
func EncryptAES(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	src = paddingText1(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src

}

// DecryptAES
// 描述：解密
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		src：待解密byte切片
//		key：密钥
//   [out]
//		解密后的byte切片
//   [in/out]
//		无
func DecryptAES(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unPaddingText1(src)
	return src
}

// AesEncrypt
// 描述：带iv向量的加密
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		src：待加密byte切片
//		key：密钥
//		iv：iv向量
//   [out]
//		加密后的byte切片
//   [in/out]
//		无
func AesEncrypt(src, key, iv []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	src = paddingText1(src, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(src, src)
	return base64.StdEncoding.EncodeToString(src)
}

// AesDecrypt
// 描述：携带iv向量的解密
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		str：待解密byte切片
//		key：密钥
//		iv：iv向量
//   [out]
//		解密后的byte切片
//   [in/out]
//		无
func AesDecrypt(src, key, iv []byte) []byte {
	src, err := base64.StdEncoding.DecodeString(string(src))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(src, src)
	src = unPaddingText1(src)

	return src
}
