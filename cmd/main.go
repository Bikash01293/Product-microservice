package main

import (
	"fmt"
	"log"
	"net"
	"product-micro/pkg/config"
	"product-micro/pkg/db"
	"product-micro/pkg/pb"
	"product-micro/pkg/services"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("unable to Load Config File")
	}

	h := db.Init(c.Dburl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	fmt.Println("Product Svc on", c.Port)

	s := services.Server { 
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}