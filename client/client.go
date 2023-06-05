package main

import (
	"context"
	"log"
	pb "medx/grpc/pb/proto"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewLoginServiceClient(conn)

	resp, err := c.DoLogin(context.Background(), &pb.LoginRequest{Auth: &pb.Auth{
		Username: "Amir",
		Password: "Mirzaei",
	},
	})

	if err != nil {
		log.Fatalf("unable to fetch data from grpc server: %v", err)
	}

	log.Printf("token is: %v", resp.Token)

}
