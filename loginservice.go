package main

import (
	"context"
	pb "medx/grpc/pb/proto"
)

func (*server) DoLogin(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{Token: "something"}, nil
}
