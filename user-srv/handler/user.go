package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	model "github.com/YuxinZhaozyx/GoMicroBookshop/user-srv/model/user"
	proto "github.com/YuxinZhaozyx/GoMicroBookshop/user-srv/proto/user"
)

type Service struct{}

var (
	userService model.Service
)

// Init 初始化handler
func Init() {
	var err error
	userService, err = model.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// QueryUserByName 通过参数中的名字返回用户
func (e *Service) QueryUserByName(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return err
	}

	rsp.User = user
	rsp.Success = true
	return nil
}
