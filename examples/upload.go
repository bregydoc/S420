package main

import (
	"context"
	"io/ioutil"

	"github.com/k0kubun/pp"

	s420con "github.com/bregydoc/S420/connection"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("18.217.206.72:24000", grpc.WithInsecure())
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

/*
docker run -p 9000:9000 --name minio1 -v /mnt/data:/data -v /mnt/config:/root/.minio -e "MINIO_ACCESS_KEY=bomboperu" -e "MINIO_SECRET_KEY=1i6sxu6pytfb391fzpl3pltg" minio/minio server /data &
*/
