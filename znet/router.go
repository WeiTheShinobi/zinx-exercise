package znet

import (
	"zinx/ziface"
)

type BasicRouter struct{}

func (b *BasicRouter) PreHandle(request ziface.IRequest) {}

func (b *BasicRouter) Handle(request ziface.IRequest) {}

func (b *BasicRouter) PostHandle(request ziface.IRequest) {}
