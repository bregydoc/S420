package main

import (
	"flag"
	"fmt"

	s420 "github.com/bregydoc/S420"
	"github.com/bregydoc/S420/backends"
	"github.com/bregydoc/S420/plugins/imods"
	s420con "github.com/bregydoc/S420/proto"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"net"
)

func main() {

	configPath := flag.String("config", "/s420.config.yml", "Where your config file is it")

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

	im := imods.NewImagePlugin()

	h := s420.NewHumanClient(mStore, conf.Public, r, im)

	go func() { h.Run() }()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Service.Port))
	if err != nil {
		logrus.Errorf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(conf.Security.ServerCertPath, conf.Security.ServerKeyPath)
	if err != nil {
		logrus.Errorf("could not load TLS keys: %s\n", err)
	}
	var grpcServer *grpc.Server
	if err == nil {
		// With credentials
		grpcServer = grpc.NewServer(grpc.Creds(creds))
		logrus.Infof("Grpc service init with credentials from: %s %s", conf.Security.ServerCertPath, conf.Security.ServerKeyPath)
	} else {
		grpcServer = grpc.NewServer()
		logrus.Info("Grpc service init without credentials")
	}

	s420con.RegisterS420Server(grpcServer, &s420.GrpcService{
		Store: mStore,
	})

	logrus.Infof("[For Developers] GRPC listening on :%d\n", conf.Service.Port)
	err = grpcServer.Serve(lis)
	if err != nil {
		logrus.Error(err)
	}
}
