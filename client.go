package s420

import (
	"fmt"
	"net/http"

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
		path := c.Param("path")
		fmt.Println(path)

		data, cType, err := client.store.GetObject(path)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}

		c.Data(http.StatusOK, string(cType), data)
		return
	})

	client.r.Run(addr)
}
