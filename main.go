package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"medx/grpc/client"
	pb "medx/grpc/pb"
	"net"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	// "google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedLoginServiceServer
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	logger, err := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logger)
	defer logger.Close()

	go runGrpcServer(&wg)

	go runGatewayServer(&wg)

	a, err := client.CallGrpcServer()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
	wg.Wait()
}

func runGrpcServer(*sync.WaitGroup) {

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterLoginServiceServer(grpcServer, &server{})

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen 9000: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}

func runGatewayServer(*sync.WaitGroup) {
	grpcMux := runtime.NewServeMux()
	err := pb.RegisterLoginServiceHandlerServer(context.Background(), grpcMux, &server{})
	if err != nil {
		log.Fatalf("fail to run grpc gateway: %v", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("failed to start listener: ", err)
	}

	err = http.Serve(listener, mux)

	if err != nil {
		log.Fatalf("cannot start http gateway server: ", err)
	}
}
