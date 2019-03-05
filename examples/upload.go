package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"io/ioutil"

	"github.com/k0kubun/pp"

	"github.com/bregydoc/S420/proto"
	"google.golang.org/grpc"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("/Users/macbook/Documents/bombo/bombo-s420-server.crt", "")
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial("resources.bombo.pe:24000", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}

	client := s420con.NewS420Client(conn)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile("/Users/macbook/KiteOnboarding.py")
	if err != nil {
		panic(err)
	}
	resp, err := client.SaveNewObject(context.TODO(), &s420con.NewObjectParams{
		Data: data,
		Options: &s420con.ObjectOptions{
			ContentType: "application/python",
		},
		Path: "bombo/kite2.py",
	})

	if err != nil {
		panic(err)
	}

	pp.Println(resp)
}
