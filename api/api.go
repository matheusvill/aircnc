package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusvill/aircnc/user"
)

func Start() {
	r := gin.Default()

	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		user := user.GetUser(id)
		json, _ := json.Marshal(use)

		c.String(http.StatusOK, json)
	})

	r.Run()
}
