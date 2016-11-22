package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matheusvill/aircnc/apartment"
	"github.com/matheusvill/aircnc/booking"
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

	r.POST("/user", func(c *gin.Context) {
		var u user.User
		c.BindJSON(&u)

		_ = user.Create(u)

		c.JSON(http.StatusOK, nil)
	})

	r.GET("/apartment/:id", func(c *gin.Context) {
		sId := c.Param("id")

		id, err := strconv.Atoi(sId)
		if err != nil {
			fmt.Println(err)
		}

		apartment := apartment.Get(id)

		c.JSON(http.StatusOK, apartment)
	})

	r.POST("/apartment", func(c *gin.Context) {
		var a apartment.Apartment
		c.BindJSON(&a)

		_ = apartment.Create(a)

		c.JSON(http.StatusOK, nil)
	})

	r.GET("/booking/:id", func(c *gin.Context) {
		sId := c.Param("id")

		id, err := strconv.Atoi(sId)
		if err != nil {
			fmt.Println(err)
		}

		booking := booking.Get(id)

		c.JSON(http.StatusOK, booking)
	})

	r.POST("/booking", func(c *gin.Context) {
		var a booking.Booking
		c.BindJSON(&a)

		_ = booking.Create(a)

		c.JSON(http.StatusOK, nil)
	})

	r.Run()
}
