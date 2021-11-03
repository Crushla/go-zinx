package zInterface

type Router interface {
	// PreHandler 处理conn业务之前的方法
	PreHandler(request Request)
	// Handler 处理conn业务的方法
	Handler(request Request)
	// PostHandler 处理conn业务之后的方法
	PostHandler(request Request)
}
