package account

import "practice/go/encrypt/server/util/encrypt"

func password(uuid, pwd string) string {
	salt := "openTheGate"
	return encrypt.MD5Sum(uuid[8:], pwd, salt)[:8]
}
