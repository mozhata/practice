package context

import "practice/go/encrypt/skeleton/reply"

type Context struct {
	Input *Param
	Resp  Response
}

type ProcessRequest func(*Context) reply.Replyer

func (c *Context) Reply() {
	c.Resp.ReplyFunc(c.Resp.ResponseWriter)
}
