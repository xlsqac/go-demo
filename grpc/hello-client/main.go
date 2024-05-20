package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "xlsqac/grpc/hello-server/proto"
)

type ClientTokenAuth struct {
}

func (c *ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "errorappId",
		"appKey": "appKey",
	}, nil
}

func (c *ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	// 连接 server
	// 启用 tls 认证
	//creds, _ := credentials.NewClientTLSFromFile("./test.pem", "*.test.com")
	//conn, err := grpc.NewClient("127.0.0.1:9091", grpc.WithTransportCredentials(creds)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(&ClientTokenAuth{}))
	conn, err := grpc.NewClient("127.0.0.1:9091", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)

	// 发送请求
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "client", Age: 18})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp.GetResponseMsg())
}
