package znet

import (
	"fmt"
	"strconv"
	"zinx/utils"
	"zinx/ziface"
)

type MsgHandler struct {
	Apis           map[uint32]ziface.IRouter
	WorkerPoolSize uint32
	TaskQueue      []chan ziface.IRequest
}

func NewMsgHandler() ziface.IMsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	router, ok := mh.Apis[request.GetMsgId()]

	if !ok {
		fmt.Println("router not register, msgId = ", request.GetMsgId())
		return
	}

	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

func (mh *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		panic("router has been exits, msgId = " + strconv.Itoa(int(msgId)))
	}

	mh.Apis[msgId] = router
	fmt.Println("add router Id = ", msgId, " success!")
}

func (mh *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)

		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

func (mh *MsgHandler) StartOneWorker(workerId int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID = ", workerId, " is Start")

	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

func (mh *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	WorkerId := request.GetConnection().GetConnId() % mh.WorkerPoolSize

	mh.TaskQueue[WorkerId] <- request
}
