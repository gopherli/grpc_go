package main

import (
	"fmt"
	"grpc_go/greeter_client/rpc"
	"grpc_go/greeter_server/proto/pb"
)

func main() {
	rpc.InitServer()
	req := new(pb.CodingRequest)
	req.Uid = 101272
	req.ActionType = pb.ActionType_Coding
	rsp, err := rpc.GetUserCodingAction(req)
	fmt.Println("err", err)
	fmt.Println("rsp", rsp)
}
