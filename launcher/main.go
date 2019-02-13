package main

import (
	s420 "github.com/bregydoc/S420"
	"github.com/bregydoc/S420/backends"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	conf, err := s420.NewConfigFromFile("s420.config.yml")
	if err != nil {
		panic(err)
	}

	mStore, err := backends.NewMinioStore(conf)
	if err != nil {
		panic(err)
	}

	h := s420.NewHumanClient(mStore, r)
	h.Run(":3300")
}
