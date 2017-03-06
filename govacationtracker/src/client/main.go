package main

import (
	"flag"
	"log"

	"fmt"

	pb "../pb"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const port = ":6000"

func main() {
	option := flag.Int("o", 1, "Command to run")
	flag.Parse()

	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewEmployeeServiceClient(conn)

	switch *option {
	case 1:
		SendMetadata(client)
	case 2:
		GetByBadgeNumber(client)
	}
}

func SendMetadata(client pb.EmployeeServiceClient) {
	md := metadata.MD{}
	md["user"] = []string{"yile"}
	md["password"] = []string{"password"}

	ctx := context.Background()
	ctx = metadata.NewContext(ctx, md)

	res, err := client.GetByBadgeNumber(ctx, &pb.GetByBadgeNumberRequest{BadgeNumber: 2010})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func GetByBadgeNumber(client pb.EmployeeServiceClient) {
	res, err := client.GetByBadgeNumber(context.Background(), &pb.GetByBadgeNumberRequest{BadgeNumber: 2010})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
