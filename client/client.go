package client

import (
	"context"
	"fmt"
	pb "medx/grpc/proto/login"

	"google.golang.org/grpc"
)

func CallGrpcServer() (string, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		return "", fmt.Errorf("did not connect to server: %v", err)
	}
	defer conn.Close()

	c := pb.NewLoginServiceClient(conn)

	resp, err := c.DoLogin(context.Background(), &pb.LoginRequest{Auth: &pb.Auth{
		Username: "Amir",
		Password: "Mirzaei",
	},
	})

	if err != nil {
		return "", fmt.Errorf("unable to fetch data from grpc server: %v", err)
	}

	return resp.Token, nil
}
