package zServer

import (
	"github.com/sirupsen/logrus"
	"go-zinx/zInterface"
	"testing"
)

type PingRouter struct {
	BaseRouter
}

func (p *PingRouter) PreHandler(request zInterface.Request) {
	logrus.Info("pre")
}

func (p *PingRouter) Handler(request zInterface.Request) {
	logrus.Info("now")
}

func (p *PingRouter) PostHandler(request zInterface.Request) {
	logrus.Info("post")
}

func TestServer_Start(t *testing.T) {
	server := NewServer("zinx")
	server.AddRouter(&PingRouter{})
	server.Run()
}
