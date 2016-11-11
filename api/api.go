package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matheusvill/aircnc/user"
)

func Start() {
	r := gin.Default()

	r.GET("/user/:id", func(c *gin.Context) {
		sId := c.Param("id")

		id, err := strconv.Atoi(sId)
		if err != nil {
			fmt.Println(err)
		}

		user := user.Get(id)

		c.JSON(http.StatusOK, user)
	})

	r.Run()
}
