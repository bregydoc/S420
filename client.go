package s420

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HumanClient describe the public client to get access to storage
type HumanClient struct {
	config Public
	store  StorageSystem
	r      *gin.Engine
}

// NewHumanClient returns a new public client
func NewHumanClient(store StorageSystem, c Public, r *gin.Engine) *HumanClient {
	return &HumanClient{
		config: c,
		r:      r,
		store:  store,
	}
}

// Run launch the http instance of the client
func (client *HumanClient) Run() {
	client.r.GET("/"+client.config.Prefix+"/*path", func(c *gin.Context) {
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

	client.r.Run(fmt.Sprintf(":%d", client.config.Port))
}
