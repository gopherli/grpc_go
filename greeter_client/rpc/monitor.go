package rpc

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"grpc_go/greeter_server/proto/pb"
)

type manager struct {
	userActionClient pb.UserActionClient
}

func newManager() *manager {
	m := new(manager)
	return m
}

var m = newManager()

func InitServer() {
	initUserActionClient()
}

func initUserActionClient() {
	var address string
	address = "127.0.0.1:10086"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("init server user_action err", err)
		return
	}
	m.userActionClient = pb.NewUserActionClient(conn)
	return
}

func GetUserCodingAction(req *pb.CodingRequest) (*pb.CodingResponse, error) {
	fmt.Println("req", req)
	rsp, err := m.userActionClient.GetUserCodingAction(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if rsp.ErrInfo != nil && rsp.ErrInfo.ErrCode != 0 && rsp.ErrInfo.ErrCode != 200 {
		return nil, errors.New(rsp.ErrInfo.ErrMessage)
	}
	return rsp, nil
}
