package main

import (
	pb "Mall/order/rpc/orderpb"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

//用于指定微服务的端口
var (
	port = flag.Int("port", 50051, "The server port") //默认设置为50051端口
)

type orderserver struct {
	pb.UnimplementedOrderServer
}

func (s *orderserver) Order(ctx context.Context, in *pb.OrdersRequest) (*pb.OrdersResponse, error) {
	log.Println("order server received")
	return &pb.OrdersResponse{}, nil
}

func main() {
	//flag包用于命令行参数的解析
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port)) //对于TPC协议，如果没有指定IP地址，那么默认监听本地所有IP地址
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() //创建一个gRPC服务器
	pb.RegisterOrderServer(s, &orderserver{})
	log.Printf("order server listening at %v", lis.Addr()) //告知order服务已经启动

	if err := s.Serve(lis); err != nil { //Serve方法会接受lis的连接请求，对于每个请求，都会创建一个新的goroutine来处理。这个goroutine会读取gRPC的请求，然后调用注册的处理函数来处理请求
		log.Fatalf("order service failed to serve: %v", err)
	}
}
