package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/bregydoc/S420"
	"github.com/bregydoc/S420/backends"
	"github.com/bregydoc/S420/connection"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"google.golang.org/grpc"
)

func main() {
	servicePort := flag.Int64("port", 24000, "Define the service port")
	configPath := flag.String("config", "s420.config.yml", "Where your config file is it")
	flag.Parse()

	r := gin.Default()
	conf, err := s420.NewConfigFromFile(*configPath)
	if err != nil {
		panic(err)
	}

	pp.Println(conf)

	mStore, err := backends.NewMinioStore(conf.Storage)
	if err != nil {
		panic(err)
	}

	h := s420.NewHumanClient(mStore, conf.Public, r)

	go func() { h.Run() }()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *servicePort))
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	s420con.RegisterS420Server(grpcServer, &s420.GrpcService{
		Store: mStore,
	})

	log.Printf("[For Developers] GRPC listening on :%d\n", *servicePort)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println(err)
	}
}
