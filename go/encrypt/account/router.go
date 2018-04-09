package account

import (
	"net/http"
	"practice/go/encrypt/skeleton/input"
	"practice/go/encrypt/skeleton/reply"
	"practice/go/encrypt/skeleton/route"
)

func NewRoute() []*route.Route {
	return []*route.Route{
		route.NewRoute(
			"/reg",
			"POST",
			reply.Wrap(register),
		),
		route.NewRoute(
			"/login",
			"POST",
			reply.Wrap(login),
		),
	}
}

func register(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	// TODO: 验证码校验
	p := struct {
		Email string `json:"email"`
		Phone string `json:"phone"`
		PWD   string `json:"pwd"`
	}{}
	err := input.NewParam(r).JSONBody(&p).Error()
	if err != nil {
		return reply.Err(err)
	}
	if err = CheckPasswordStrength(p.PWD); err != nil {
		return reply.Err(err)
	}

	// TODO: 校验邮箱/手机是否已经注册
	// 邮箱注册
	var uid string
	if p.Email != "" {
		uid, err = regByEmail(p.Email, p.PWD)
		if err != nil {
			return reply.Err(err)
		}
	}
	// 手机注册
	uid, err = regByPhone(p.Phone, p.PWD)
	if err != nil {
		return reply.Err(err)
	}

	// TODO: 注册成功自动登录
	return reply.Success(map[string]interface{}{
		"user": uid,
	})
}
func login(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	p := struct {
		Email string `json:"email"`
		PWD   string `json:"pwd"`
	}{}
	if err := input.NewParam(r).JSONBody(&p).Error(); err != nil {
		return reply.Err(err)
	}

	// TODO: 校验参数
	user, err := loginByEmail(p.Email, p.PWD)
	if err != nil {
		return reply.Err(err)
	}
	// TODO: token

	// TODO: {code: xx, msg: xx, body: obj}
	return reply.Success(map[string]interface{}{
		"user": user,
	})
}
