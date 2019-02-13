package s420

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HumanClient struct {
	store Storage
	r     *gin.Engine
}

func NewHumanClient(store Storage, r *gin.Engine) *HumanClient {
	return &HumanClient{
		r:     r,
		store: store,
	}
}

func (client *HumanClient) Run(addr string) {
	client.r.GET("/public/*path", func(c *gin.Context) {
		fmt.Println(c.Param("path"))
	})

	client.r.Run(addr)
}
