package account

import (
	"database/sql"

	"github.com/mozhata/merr"
)

type User struct {
	ID    int    `json:"-"`
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Bio   string `json:"bio"`
}

func (u *User) isValid() bool {
	return u.UUID != "" &&
		u.Name != "" &&
		(u.Email != "" || u.Phone != "")
}

func CreateUser(db *sql.DB, u User) error {
	sql_ := "insert into user(uuid, name, phone, email, bio) value (?, ?, ?, ?, ?)"
	_, err := db.Exec(sql_, u.UUID, u.Name, u.Phone, u.Email, u.Bio)
	if err != nil {
		return merr.WrapErr(err)
	}
	return nil
}

func UserByUUID(db *sql.DB, uuid string) (*User, error) {
	sql_ := "select uuid, name, phone, email, bio from user where uuid = ?"
	var u User
	err := db.QueryRow(sql_, uuid).Scan(&u.UUID, &u.Name, &u.Phone, &u.Email, &u.Bio)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, merr.WrapErr(nil, "user not found by uuid %s", uuid)
		}
		return nil, merr.WrapErr(err)
	}
	return &u, nil
}
