package handler

import (
	"context"
	"fmt"

	"github.com/davveo/tsquare/srv/uuid/idworker"

	uuidproto "github.com/davveo/tsquare/proto/uuid"
)

type Uuid struct{}

func (e *Uuid) GenerateId(ctx context.Context, req *uuidproto.Request, rsp *uuidproto.Response) error {
	if err := idworker.InitNode(req.NodeId); err != nil {
		rsp.Msg = fmt.Sprintf(err.Error())
		rsp.Code = -1
		return nil
	}

	var idGen *idworker.IdWorker
	for i := 0; i < 5; i++ {
		idGen = idworker.GetIdWorker()
		if idGen != nil {
			break
		}
	}

	if idGen == nil {
		rsp.Msg = fmt.Sprintf("IdWorker对象初始化失败，请重试!")
		rsp.Code = -1
		return nil
	}

	id, err := idGen.GetNextId()

	if err != nil {
		rsp.Msg = fmt.Sprintf("Id生成失败!")
		rsp.Code = -1
		return nil
	}

	rsp.Data = &uuidproto.Data{
		Id: id,
	}
	rsp.Msg = "success"
	rsp.Code = 0

	return nil
}
