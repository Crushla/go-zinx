package zInterface

type MsgHandler interface {
	// DoMsgHandler 调度Router消息处理方法
	DoMsgHandler(request Request)
	// AddRouter 为消息添加具体的逻辑
	AddRouter(msgID uint32, router Router)
	StartWorkerPool()
	StartOneWorker()
	SendMsgToTaskQueue(request Request)
}
