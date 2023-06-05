package main

import (
	"context"
	"log"
	"net/http"

	pb "medx/grpc/proto/login"
	"net"
	"os"
	"runtime"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedLoginServiceServer
}

func main() {

	logger, err := os.OpenFile("app.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logger)
	defer logger.Close()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	pb.RegisterLoginServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:9000",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterLoginServiceHandler(
		context.Background(),
		gwmux,
		conn,
	)

	if err != nil {
		log.Fatalln("fail to register gateway", err)
	}

	gwServer := &http.Server{
		Addr:    ":8001",
		Handler: gwmux,
	}
	log.Println("Serving GRPC-Gateway on port 8000")
	log.Fatalln(gwServer.ListenAndServe())

}
