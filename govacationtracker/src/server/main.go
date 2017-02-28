package main

import (
    "log"
    "net"
    pb "../pb"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const port = ":6000"

type employeeService struct{}

func(s *employeeService) GetByBadgeNumber(ctx context.Context, req *pb.GetByBadgeNumberRequest) (*pb.EmployeeResponse, error){
    return nil, nil
}

func(s *employeeService) GetAll(req *pb.GetAllRequest, stream pb.EmployeeService_GetAllServer) (error){
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