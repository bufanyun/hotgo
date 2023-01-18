// Package encrypt
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package encrypt

import (
	"encoding/base64"
	"github.com/forgoer/openssl"
)

// AesECBEncrypt 加密
func AesECBEncrypt(src, key []byte) (dst []byte, err error) {
	return openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
}

// AesECBDecrypt 解密
func AesECBDecrypt(src, key []byte) (dst []byte, err error) {
	return openssl.AesECBDecrypt(src, key, openssl.PKCS7_PADDING)
}

// MustAesECBEncryptToString
// Return the encryption result directly. Panic error
func MustAesECBEncryptToString(bytCipher, key string) string {
	dst, err := AesECBEncrypt([]byte(bytCipher), []byte(key))
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(dst)
}

// MustAesECBDecryptToString
// Directly return decryption result, panic error
func MustAesECBDecryptToString(bytCipher, key string) string {
	dst, err := AesECBDecrypt([]byte(bytCipher), []byte(key))
	if err != nil {
		panic(err)
	}
	return string(dst)
}
