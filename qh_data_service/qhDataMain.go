package main

import (
	"fmt"
	"log"
	"net"
	qhdata_proto "qhdataservice/proto"
	"qhdataservice/qhdata"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("[main] qh data service")

	svc, err := qhdata.NewService()
	if err != nil {
		fmt.Println("[main] failed to create service: ", err)
	}
	startGrpcServer(&svc)
}

func startGrpcServer(svc *qhdata.Service) {
	grpcPort := 50099
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Transport [tcp] Listening on: %d", grpcPort)

	grpcServer := grpc.NewServer()

	fmt.Println("[main] Service object created")

	qhdata_proto.RegisterQHDataServer(grpcServer, svc)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
