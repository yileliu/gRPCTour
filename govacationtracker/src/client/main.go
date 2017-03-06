package main

import (
	"flag"
	"io"
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
	case 3:
		GetAll(client)
	case 4:
		SaveAll(client)
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

func GetAll(client pb.EmployeeServiceClient) {
	stream, err := client.GetAll(context.Background(), &pb.GetAllRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		emp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(emp.Employee)
	}
}

func SaveAll(client pb.EmployeeServiceClient) {
	newEmployees := []pb.Employee{
		pb.Employee{
			BadgeNumber:         229,
			FirstName:           "Amity",
			LastName:            "Fuller",
			VacationAccrualRate: 2.3,
			VacationAccrued:     23.4,
		},
		pb.Employee{
			BadgeNumber:         230,
			FirstName:           "Amity",
			LastName:            "Fuller",
			VacationAccrualRate: 2.3,
			VacationAccrued:     23.4,
		},
	}

	stream, err := client.SaveAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	doneCh := make(chan struct{})

	go (func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				doneCh <- struct{}{}
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(res.Employee)
		}
	})()

	for _, e := range newEmployees {
		err := stream.Send(&pb.EmployeeRequest{Employee: &e})

		if err != nil {
			log.Fatal(err)
		}
	}

	stream.CloseSend()

	<-doneCh
}
