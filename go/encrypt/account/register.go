package account

import (
	"github.com/pborman/uuid"

	"practice/go/encrypt/db"
)

func regByEmail(username, email, pwd string) (*User, error) {
	// TODO: check wether email has registered
	u := User{
		UUID:  uuid.New(),
		Name:  username,
		Email: email,
	}

	err := CreateUser(db.MySQL, u)
	if err != nil {
		return nil, err
	}

	localAuth := NewLocalAuth(u.UUID, email, "", pwd)
	err = CreateLocalAuth(db.MySQL, localAuth)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
