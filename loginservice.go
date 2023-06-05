package main

import (
	"context"
	pb "medx/grpc/proto/login"
)

func (*server) DoLogin(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{Token: "something"}, nil
}
