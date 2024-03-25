package main

import (
	"context"
	"fmt"
	//"grpc/pb"
	pb "grpc_service/pb_gen/grpc"
	"net"

	"google.golang.org/grpc"
)

// hello server

type server struct {
	pb.GrpcSerivceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.GrpcRequest) (*pb.GrpcResponse, error) {
	return &pb.GrpcResponse{Reply: "Hello " + in.Name + ", welcome to grpc"}, nil
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
