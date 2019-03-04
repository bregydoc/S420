package main

import (
	"context"
	"io/ioutil"

	"github.com/k0kubun/pp"

	"github.com/bregydoc/S420/connection"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:24000", grpc.WithInsecure())
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
		Path: "bombo/kite.py",
	})

	if err != nil {
		panic(err)
	}

	pp.Println(resp)
}
