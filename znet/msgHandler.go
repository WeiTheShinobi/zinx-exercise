package znet

import (
	"fmt"
	"strconv"
	"zinx/ziface"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
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
