package zutils

import (
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"go-zinx/zInterface"
	"reflect"
)

// GlobalObj 存储一起有关框架的全局参数
type GlobalObj struct {
	//Server
	TcpServer zInterface.Server `ini:"TcpServer"`
	Host      string            `ini:"Host"`
	Port      int               `ini:"Port"`
	Name      string            `ini:"Name"`
	//Zinx
	Version           string `ini:"Version"`
	MaxConn           int    `ini:"MaxConn"`
	MaxPackageSize    uint32 `ini:"MaxPackageSize"`
	WorkerPoolSize    uint32 `ini:"WorkerPoolSize"`
	MaxWorkerTaskSize uint32 `ini:"MaxWorkerTaskSize"`
}

// GlobalObject 定义一个全局对外Globalobj
var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	var configObject = new(GlobalObj)
	//读配置文件
	error := ini.MapTo(configObject, "../zconf/conf.ini")
	if error != nil {
		logrus.Error("load config failed, err", error)
		return
	}
	val := reflect.ValueOf(configObject).Elem()
	vType := val.Type()
	gval := reflect.ValueOf(GlobalObject).Elem()
	for i := 0; i < val.NumField(); i++ {
		name := vType.Field(i).Name
		if ok := gval.FieldByName(name).IsValid() && !val.Field(i).IsZero(); ok {
			gval.FieldByName(name).Set(reflect.ValueOf(val.Field(i).Interface()))
		}
	}
}

//初始化
func init() {
	GlobalObject = &GlobalObj{
		Name:              "zinx",
		Version:           "v1",
		Port:              9090,
		Host:              "127.0.0.1",
		MaxConn:           1000,
		MaxPackageSize:    4096,
		WorkerPoolSize:    10,
		MaxWorkerTaskSize: 5000,
	}
	//尝试配置文件中加载
	GlobalObject.Reload()
}
