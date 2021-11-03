package zInterface

type Server interface {
	// Start 启动服务器
	Start()
	// Stop 关闭服务器
	Stop()
	// Run 运行服务器
	Run()
	// AddRouter 路由功能
	AddRouter(router Router)
}
