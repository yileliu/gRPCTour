package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	pb "../pb"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const port = ":6000"

var nextId int32 = 100

type employeeService struct{}

func (s *employeeService) GetByBadgeNumber(ctx context.Context, req *pb.GetByBadgeNumberRequest) (*pb.EmployeeResponse, error) {
	if md, ok := metadata.FromContext(ctx); ok {
		fmt.Println("Metadata received: %v\n", md)
	}

	for _, e := range employees {
		if req.BadgeNumber == e.BadgeNumber {
			fmt.Println("Employee found: %v\n", e)
			return &pb.EmployeeResponse{Employee: &e}, nil
		}
	}

	return nil, errors.New("Employee not found")
}

func (s *employeeService) GetAll(req *pb.GetAllRequest, stream pb.EmployeeService_GetAllServer) error {
	for _, e := range employees {
		stream.Send(&pb.EmployeeResponse{Employee: &e})
	}

	return nil
}

// func(s *employeeService) AddPhoto(stream pb.EmployeeService_AddPhotoServer) (error){
//     if md, ok := metadata.FromContext(ctx); ok{
//         fmt.Println("Receiveing photo for badge number: %v\n", md["badgenumber"][0])
//     }

//     imgData := []byte{}
//     for{
//         data, err := stream.Recv()
//         if err == io.EOF{
//             fmt.Println("File received with lenght %v\n", len(imgData))
//             return stream.SendAndClose()
//         }

//         if err != nil {
//             return err
//         }
//         imgData = append(imgData, data)
//     }

//     return nil
// }

func GetNextEmployeeId() {
	nextId = nextId + 1
	fmt.Println(nextId)
}

func (s *employeeService) Save(ctx context.Context, req *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	employee := req.Employee

	employees = append(employees, *employee)

	return &pb.EmployeeResponse{Employee: employee}, nil
}
func (s *employeeService) SaveAll(stream pb.EmployeeService_SaveAllServer) error {
	for {
		var err error
		var emp *pb.EmployeeRequest
		emp, err = stream.Recv()

		if emp.Employee.Id == 0 {
			GetNextEmployeeId()
			emp.Employee.Id = nextId
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fmt.Println(emp.Employee)
		//employees = append(employees, *emp.Employee)
		err = stream.Send(&pb.EmployeeResponse{Employee: emp.Employee})

		if err != nil {
			return err
		}
	}

	fmt.Println("Here")

	fmt.Println(len(employees))
	for _, e := range employees {
		fmt.Println(e)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterEmployeeServiceServer(s, &employeeService{})
	log.Println("Starting server on port" + port)
	s.Serve(lis)
}
