package zServer

import (
	"github.com/sirupsen/logrus"
	"go-zinx/zInterface"
	"go-zinx/zutils"
	"strconv"
)

type MsgHandler struct {
	Apis           map[uint32]zInterface.Router
	TaskQueue      []chan Request
	WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]zInterface.Router),
		TaskQueue:      make([]chan Request, zutils.GlobalObject.WorkerPoolSize),
		WorkerPoolSize: zutils.GlobalObject.WorkerPoolSize,
	}
}

func (m *MsgHandler) DoMsgHandler(request zInterface.Request) {
	router, ok := m.Apis[request.GetMsgID()]
	if !ok {
		logrus.Error("api msgID ", request.GetMsgID(), "is not found")
	}
	router.PreHandler(request)
	router.Handler(request)
	router.PostHandler(request)
}

func (m *MsgHandler) AddRouter(msgID uint32, router zInterface.Router) {
	//判断当前msg绑定的API是否存在
	if _, ok := m.Apis[msgID]; ok {
		panic("repeat api, msgID =" + strconv.Itoa(int(msgID)))
	}
	//添加存在msg与API绑定关系
	m.Apis[msgID] = router
	logrus.Info("add router success")
}

func (m *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		m.TaskQueue[i] = make(chan Request, zutils.GlobalObject.MaxWorkerTaskSize)
		go m.StartOneWorker(i, m.TaskQueue[i])
	}
}

func (m *MsgHandler) StartOneWorker(workerID int, taskQueue chan Request) {
	logrus.Info("worker id ", workerID)
	for true {
		select {
		case request := <-taskQueue:
			m.DoMsgHandler(&request)
		}
	}
}

func (m *MsgHandler) SendMsgToTaskQueue(request Request) {
	workerId := request.GetConnection().GetConnID() % m.WorkerPoolSize
	m.TaskQueue[workerId] <- request
}
