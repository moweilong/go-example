package simple

import (
	"hotgo/internal/consts"
	"hotgo/utility/encrypt"

	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/errors/gerror"
)

// DecryptText 解密文本
func DecryptText(text string) (string, error) {
	str, err := gbase64.Decode([]byte(text))
	if err != nil {
		return "", err
	}

	str, err = encrypt.AesECBDecrypt(str, consts.RequestEncryptKey)
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
