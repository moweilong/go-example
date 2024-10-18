package main

import (
	"context"
	"os"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var RequestEncryptKey = []byte("f080a463654b2279")

func main() {
	ctx := context.Background()
	dst, err := EncryptText("123456")
	if err != nil {
		g.Log().Error(ctx, "加密文本错误", err)
		os.Exit(1)
	}

	bDst := gbase64.Encode([]byte(dst))
	g.Log().Info(ctx, "解密文本", string(bDst))

	src, err := DecryptText(string(bDst))
	if err != nil {
		g.Log().Error(ctx, "解密文本错误", err)
		os.Exit(1)
	}
	g.Log().Info(ctx, "解密文本", src)
}

// EncryptText 加密文本
func EncryptText(text string) (string, error) {
	src := []byte(text)
	dst, err := AesECBEncrypt(src, RequestEncryptKey)
	if err != nil {
		return "", err
	}
	return string(dst), nil
}

// DecryptText 解密文本
func DecryptText(text string) (string, error) {
	str, err := gbase64.Decode([]byte(text))
	if err != nil {
		return "", err
	}

	str, err = AesECBDecrypt(str, RequestEncryptKey)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

// CheckPassword 检查密码
func CheckPassword(input, salt, hash string) (err error) {
	// 解密密码
	password, err := DecryptText(input)
	if err != nil {
		return err
	}

	if hash != gmd5.MustEncryptString(password+salt) {
		err = gerror.New("用户密码不正确")
		return
	}
	return
}
