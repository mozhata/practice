package account

import "practice/go/encrypt/db"

func loginByEmail(email, pwd string) (*User, error) {
	uuid, err := CheckAuthByEmail(email, pwd)
	if err != nil {
		return nil, err
	}
	user, err := UserByUUID(db.MySQL, uuid)
	if err != nil {
		return nil, err
	}
	return user, nil

}
