// Copyright (c) 2023 YaoBase
// 描述: 定义rsa加密相关加解密
// 创建者: 张平
// 创建时间: 2022/5/26 17:13

package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path"
)

// MakeKeyPair
// 描述：制作RSA密钥
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		outPath：输出的目录
//		bits：密钥大小
//   [out]
//		制作RSA密钥过程中出现的错误
//   [in/out]
//		无
func MakeKeyPair(outPath string, bits int) error {
	priKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	x509PriKey := x509.MarshalPKCS1PrivateKey(priKey)
	priFile, err := os.Create(path.Join(outPath, "private.key"))
	if err != nil {
		return err
	}
	defer priFile.Close()

	perBlock := pem.Block{Type: "PUBLIC KEY", Bytes: x509PriKey}
	if err := pem.Encode(priFile, &perBlock); err != nil {
		return err
	}

	pubKey := priKey.PublicKey
	x509PubKey, _ := x509.MarshalPKIXPublicKey(&pubKey)
	pemPubKey := pem.Block{Type: "PRIVATE KEY", Bytes: x509PubKey}
	pubFile, _ := os.Create(path.Join(outPath, "public.key"))
	defer pubFile.Close()
	if err := pem.Encode(pubFile, &pemPubKey); err != nil {
		return err
	}

	return nil
}

// RSAEncrypt
// 描述：rsa加密
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		pubKey：公钥
//		data：待加密的数据
//   [out]
//		结果1：加密后的数据
//		结果2：加密过程中出现的错误
//   [in/out]
//		无
func RSAEncrypt(pubKey, data []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), data)
}

// RSADecrypt
// 描述：rsa解密
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		priKey：私钥
//		data：待解密的数据
//   [out]
//		结果1：解密后的数据
//		结果2：解密过程中出现的错误
//   [in/out]
//		无
func RSADecrypt(priKey, data []byte) ([]byte, error) {
	block, _ := pem.Decode(priKey)
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, pri, data)
}

// RSASign
// 描述：获取rsa签名
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		priKey：私钥
//		salt：盐值
//		data：获取rsa签名原数据
//   [out]
//		结果1：rsa签名数据
//		结果2：获取rsa签名过程中出现的错误
//   [in/out]
//		无
func RSASign(priKey, salt, data []byte) ([]byte, error) {
	block, _ := pem.Decode(priKey)
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(data)
	hash = sha256.Sum256(append(salt, hash[:]...))
	return rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hash[:])
}

// RSAVerify
// 描述：rsa签名校验
// 创建者：张平
// 创建时间：2022/5/26 17:23
// 参数：
//	 [in]
//		pubKey：公钥
//		sign：待校验的签名数据
//		salt：盐值
//		data：rsa签名原数据
//   [out]
//		获取rsa签名过程中出现的错误
//   [in/out]
//		无
func RSAVerify(pubKey, sign, salt, data []byte) error {
	block, _ := pem.Decode(pubKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	hash := sha256.Sum256(data)
	hash = sha256.Sum256(append(salt, hash[:]...))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hash[:], sign)
}
