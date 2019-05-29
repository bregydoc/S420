package s420

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HumanClient describe the public client to get access to storage
type HumanClient struct {
	config  Public
	store   StorageSystem
	r       *gin.Engine
	Plugins []Plugin
}

// NewHumanClient returns a new public client
func NewHumanClient(store StorageSystem, c Public, r *gin.Engine, plugins ...Plugin) *HumanClient {
	return &HumanClient{
		config:  c,
		r:       r,
		store:   store,
		Plugins: plugins,
	}
}

// Run launchs the http instance of the client
func (client *HumanClient) Run() {
	client.r.GET("/"+client.config.Prefix+"/*path", func(c *gin.Context) {
		path := c.Param("path")
		fmt.Println(path)

		data, cType, err := client.store.GetObject(path)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		// Applying plugins
		out := data
		for _, p := range client.Plugins {
			out = p.ProcessFile(c.Request, cType, data)
		}

		c.Data(http.StatusOK, string(cType), out)
		return
	})

	client.r.Run(fmt.Sprintf(":%d", client.config.Port))
}
