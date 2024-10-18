package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"hash/fnv"

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

// Md5ToString 生成md5
func Md5ToString(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// Md5 生成md5
func Md5(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func Hash32(b []byte) uint32 {
	h := fnv.New32a()
	h.Write(b)
	return h.Sum32()
}
