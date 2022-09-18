package main

import (
	pb "Mall/product/rpc/productpb"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

//用于指定微服务的端口
var (
	port = flag.Int("port", 50052, "The server port") //默认设置为50052端口
)

type productserver struct {
	pb.UnimplementedProductServer
}

//实现rpc service
func (s *productserver) Products(ctx context.Context, in *pb.ProductRequest) (*pb.ProductResponse, error) {
	log.Println("product server received")
	return &pb.ProductResponse{}, nil
}

func main() {
	//flag包用于命令行参数的解析
	flag.Parse()
	//创建TCP socket
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port)) //对于TPC协议，如果没有指定IP地址，那么默认监听本地所有IP地址
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//创建gRPC服务器
	s := grpc.NewServer()
	pb.RegisterProductServer(s, &productserver{})
	log.Printf("product server listening at %v", lis.Addr()) //告知product服务已经启动

	if err := s.Serve(lis); err != nil {
		log.Fatalf("product service failed to serve: %v", err)
	}
}
