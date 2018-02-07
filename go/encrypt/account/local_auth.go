package account

import (
	"database/sql"
	"practice/go/encrypt/db"
	"practice/go/encrypt/skeleton/common"

	"github.com/mozhata/merr"
)

type LocalAuth struct {
	UUID     string
	Email    string
	Phone    string
	Password string
}

func NewLocalAuth(uuid, email, phone, pwd string) *LocalAuth {
	return &LocalAuth{
		UUID:     uuid,
		Email:    email,
		Phone:    phone,
		Password: password(uuid, pwd),
	}
}

func (l *LocalAuth) isValid() bool {
	return l.UUID != "" &&
		l.Password != "" &&
		(l.Email != "" || l.Phone != "")
}

func CreateLocalAuth(db *sql.DB, auth *LocalAuth) error {
	sql := "insert into local_auth (uuid, email, phone, password) value (?, ?, ?, ?)"
	_, err := db.Exec(sql, auth.UUID, auth.Email, auth.Phone, auth.Password)
	if err != nil {
		return merr.WrapErr(err)
	}
	return nil
}

func GetLocalAuthByEmail(db *sql.DB, email string) (*LocalAuth, error) {
	sql_ := "select * from local_auth where email = ?"

	l := LocalAuth{}
	err := db.QueryRow(sql_, email).Scan(
		&l.UUID,
		&l.Email,
		&l.Phone,
		&l.Password,
	)
	if err == sql.ErrNoRows {
		return nil, merr.WrapErr(nil, "local auth not found by email %s", email)
	}
	if err != nil {
		return nil, merr.WrapErr(err)
	}

	return &l, nil
}

func CheckAuthByEmail(email, pwd string) (uuid string, err error) {
	localAuth, err := GetLocalAuthByEmail(db.MySQL, email)
	if err != nil {
		return "", err
	}
	if password(localAuth.UUID, pwd) == localAuth.Password {
		return localAuth.UUID, nil
	}
	return "", common.InvalidArgumentErr("invalid password")
}
