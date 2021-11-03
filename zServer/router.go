package zServer

import (
	"go-zinx/zInterface"
)

//实现router时，嵌入这个基类，然后根据需要对这个基类方法进行重写
type BaseRouter struct{}

func (b *BaseRouter) PreHandler(request zInterface.Request) {}

func (b *BaseRouter) Handler(request zInterface.Request) {}

func (b *BaseRouter) PostHandler(request zInterface.Request) {}
