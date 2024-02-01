package commonFunc

import "github.com/gogf/gf/v2/crypto/gmd5"

// EncryptPassword 密码加密
func EncryptPassword(password string) (newPsw string) {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password))
}
