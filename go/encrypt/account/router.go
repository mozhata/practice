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

func register(w http.ResponseWriter, r *http.Request) reply.Replyer {
	p := struct {
		Email string `json:"email"`
		PWD   string `json:"pwd"`
		User  string `json:"user"`
	}{}
	if err := input.NewParam(r).JSONBody(&p).Error(); err != nil {
		return reply.Err(err)
	}

	// TODO: 校验参数
	// TODO: 校验邮箱/手机是否已经注册
	user, err := regByEmail(p.User, p.Email, p.PWD)
	if err != nil {
		return reply.Err(err)
	}
	return reply.JSON(map[string]interface{}{
		"user": user,
	})
}
func login(w http.ResponseWriter, r *http.Request) reply.Replyer {
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
	return reply.JSON(map[string]interface{}{
		"user": user,
	})
}
