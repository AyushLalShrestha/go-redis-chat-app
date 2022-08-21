package message

import "github.com/AyushLalShrestha/go-redis-chat-app/rediscli"

type Controller struct {
	r *rediscli.Redis
}

func NewController(r *rediscli.Redis) *Controller {
	return &Controller{
		r: r,
	}
}
