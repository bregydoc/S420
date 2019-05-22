package main

import (
	"context"
	"io/ioutil"
	"log"

	s420con "github.com/bregydoc/S420/proto"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("./server.crt", "")
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial("localhost:24000", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}

	client := s420con.NewS420Client(conn)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadFile("./ghost.jpg")
	if err != nil {
		panic(err)
	}
	resp, err := client.SaveNewObject(context.TODO(), &s420con.NewObjectParams{
		Data: data,
		Options: &s420con.ObjectOptions{
			ContentType: "image/jpg",
		},
		Path: "covers/ghost.jpg",
	})

	if err != nil {
		panic(err)
	}

	// You can access to localhost:3300/public/covers/ghost.jpg

	log.Println(resp)
}
