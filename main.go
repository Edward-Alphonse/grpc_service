package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Edward-Alphonse/grpc_service/pb_gen/grpc"
	"google.golang.org/grpc"
)

// hello server

type server struct {
	pb.GrpcSerivceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.GrpcRequest) (*pb.GrpcResponse, error) {
	log.Printf("SayHello: %v", in)
	resp := new(pb.GrpcResponse)
	resp.BaseResponse = &pb.BaseResponse{
		Code:    200,
		Message: "success",
	}
	resp.Reply = fmt.Sprintf("hello, %s", in.Name)
	return resp, nil
}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                        // 创建gRPC服务器
	pb.RegisterGrpcSerivceServer(s, new(server)) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
