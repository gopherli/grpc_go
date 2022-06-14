package rpc

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_go/greeter_server/proto/pb"
	"log"
	"net"
)

type Monitor struct {
	pb.UnimplementedUserActionServer
}

func newMonitor() *Monitor {
	m := new(Monitor)
	return m
}

func Start() {
	m := newMonitor()
	lis, err := net.Listen("tcp", "localhost:10086")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserActionServer(s, m)
	log.Printf("server listening at %v", lis.Addr())
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (m *Monitor) GetUserCodingAction(ctx context.Context, req *pb.CodingRequest) (*pb.CodingResponse, error) {
	rsp, err := GetUserCodingActionRsp(req)
	if err != nil {
		rsp = new(pb.CodingResponse)
		rsp.ErrInfo = new(pb.ErrInfo)
		rsp.ErrInfo.ErrCode = 500
		rsp.ErrInfo.ErrMessage = err.Error()
		return rsp, nil
	}
	return rsp, nil
}

func GetUserCodingActionRsp(req *pb.CodingRequest) (*pb.CodingResponse, error) {
	rsp := new(pb.CodingResponse)
	userBaseInfo := new(pb.UserBaseInfo)
	if req.Uid != 101272 {
		rsp.ErrInfo.ErrCode = 500
		rsp.ErrInfo.ErrMessage = fmt.Sprintf("当前用户%d不存在", req.Uid)
		return rsp, errors.New("not nil")
	}
	if req.Uid == 101272 {
		userBaseInfo.Uid = req.Uid
		userBaseInfo.Name = "李二蛋"
		userBaseInfo.Age = 1997
		userBaseInfo.Home = "湖南株洲"
		userBaseInfo.Sex = 1
		rsp.UserBaseInfo = userBaseInfo
	}
	if req.ActionType == pb.ActionType_Coding {
		rsp.Language = "Go"
	}
	return rsp, nil
}
