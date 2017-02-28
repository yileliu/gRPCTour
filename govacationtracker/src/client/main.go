package main

import (
    "log"
	grpc "google.golang.org/grpc" 
    pb "../pb"
)

const port = ":6000"

func main(){
    conn, err := grpc.Dial("localhost" + port, grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
    }

    defer conn.Close()

    pb.NewEmployeeServiceClient(conn)
}