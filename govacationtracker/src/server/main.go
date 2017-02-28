package main

import (
    "log"
    "net"
    "errors"
    "fmt"
    pb "../pb"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
    "google.golang.org/grpc/metadata"
)

const port = ":6000"

type employeeService struct{}

func(s *employeeService) GetByBadgeNumber(ctx context.Context, req *pb.GetByBadgeNumberRequest) (*pb.EmployeeResponse, error){
    if md, ok := metadata.FromContext(ctx); ok{
        fmt.Println("Metadata received: %v\n", md)
    }

    for _, e := range employees{
        if(req.BadgeNumber == e.BadgeNumber){
            return &pb.EmployeeResponse{Employee: &e}, nil
        }
    }

    return nil, errors.New("Employee not found")
}

func(s *employeeService) GetAll(req *pb.GetAllRequest, stream pb.EmployeeService_GetAllServer) (error){
    for _, e := range employees{
        stream.Send(&pb.EmployeeResponse{Employee: &e})
    }

    return nil
}

func (s *employeeService) Save(ctx context.Context, req *pb.EmployeeRequest) (*pb.EmployeeResponse, error){
    return nil, nil
}
func (s *employeeService) SaveAll(stream pb.EmployeeService_SaveAllServer) error{
    return nil
}

func main(){
    lis, err := net.Listen("tcp", port)

    if err != nil{
        log.Fatal(err)
    }
  
	s := grpc.NewServer()

	pb.RegisterEmployeeServiceServer(s, &employeeService{})
    log.Println("Starting server on port" + port)
	s.Serve(lis)
}