package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/services"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("failed to listen: %v", err)
		return
	}
	rpcServer := grpc.NewServer()                                          // 创建gRPC服务器
	services.RegisterProdServiceServer(rpcServer, &services.ProdService{}) // 在gRPC服务端注册服务

	err = rpcServer.Serve(lis)
	if err != nil {
		fmt.Println("failed to server: %v", err)
		return
	}

	//http
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	rpcServer.ServeHTTP(writer, request)
	//})
	//httpServer := &http.Server{
	//	Addr:    "8081",
	//	Handler: mux,
	//}
	//httpServer.ListenAndServe()

}
