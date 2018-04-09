package account

import (
	"github.com/pborman/uuid"

	"practice/go/encrypt/db"
)

func regByEmail(email, pwd string) (string, error) {
	uid := uuid.New()
	if err := CreateLocalAuth(db.MySQL, uid, email, "", pwd); err != nil {
		return "", err
	}
	return uid, nil
}

func regByPhone(phone, pwd string) (string, error) {
	uid := uuid.New()
	if err := CreateLocalAuth(db.MySQL, uid, "", phone, pwd); err != nil {
		return "", err
	}
	return uid, nil
}
