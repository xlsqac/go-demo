package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	pb "xlsqac/grpc/hello-server/proto"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("not token")
	}
	var appId string
	var appKey string

	if val, ok := md["appid"]; ok {
		appId = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appKey = val[0]
	}

	if appId == "appId" && appKey == "appKey" {
		respMsg := "Hello" + req.RequestName
		log.Println(respMsg)
		return &pb.HelloResponse{ResponseMsg: respMsg}, nil
	}
	return nil, errors.New("error token")
}

func main() {
	// 监听端口
	// 启动 tls 认证
	//creds, _ := credentials.NewServerTLSFromFile("../cert/server.pem", "../cert/server.key")
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// 创建服务
	//grpcServer := grpc.NewServer(grpc.Creds(creds)) // 启用 tls 认证
	grpcServer := grpc.NewServer()
	// 绑定服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("Failed to start ", err)
		return
	}
}
