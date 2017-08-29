package controllers

import (
	"encoding/json"
	"time"

	"github.com/mozhata/merr"

	"practice/go/beedemo/common"
	"practice/go/beedemo/controllers/basecontroller"
	"practice/go/beedemo/models"
	"practice/go/beedemo/module/user"
)

// UserController Operations about Users
type UserController struct {
	basecontroller.Controller
}

// Search search users by name, if filter is emtpy, return all
// @router / [get]
func (c *UserController) Search() {
	// TODO: build query struct
	// just list all user for now
	all, err := user.AllUsers()
	if err != nil {
		c.HandleErr(err)
		return
	}
	c.Success(common.M{
		"users": all,
	})
}

// CreateUser create user
// @router / [post]
func (c *UserController) CreateUser() {
	var u models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &u)
	if err != nil {
		err = merr.InvalidErr(err, "reurest body is not valid")
		c.HandleErr(err)
		return
	}
	now := time.Now()
	u.CreateTime = now
	u.UpdateTime = now

	if !u.IsValid() {
		err = merr.InvalidErr(nil, "user %#v is not valid", u)
		c.HandleErr(err)
		return
	}
	uid, err := user.CreateUser(u)
	if err != nil {
		err = merr.InternalError(err, "blablsss%s", "sdf")
		c.HandleErr(err)
		return
	}
	c.Success(common.M{
		"uid": uid,
	})
}

// @router /:uid [delete]
func (c *UserController) DeleteUser() {
	uid, err := c.GetUint64(":uid")
	if err != nil {
		err = merr.InvalidErr(err, "param uid is not vaid")
		c.HandleErr(err)
		return
	}
	err = user.DeleteUser(uid)
	if err != nil {
		c.HandleErr(err)
		return
	}
	c.Success(nil)
}

// CheckExistence check whether the given name already exist
// @router /:uname/existance [get]
func (c *UserController) CheckExistence() {
	uname := c.GetString(":uname")
	exist, err := user.CheckExistance(uname)
	if err != nil {
		c.HandleErr(err)
		return
	}
	c.Success(common.M{
		"exist": exist,
	})
}
