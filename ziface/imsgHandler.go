package ziface

type IMsgHandler interface {
	DoMsgHandler(request IRequest)
	AddRouter(msgId uint32, router IRouter)
	StartWorkerPool()
	StartOneWorker(workerId int, taskQueue chan IRequest)
	SendMsgToTaskQueue(request IRequest)
}
