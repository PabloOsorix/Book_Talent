package main

import (
	"log"

	"net/http"

	"github.com/PabloOsorix/Book_Talent/engine"
	"github.com/gin-gonic/gin"
)

type User = engine.User

func main() {

	client, err := engine.Create()
	if err != nil {
		log.Fatal(err)
	}
	defer engine.Disconnect(client)

	usersColl, err := engine.Collection(client)
	if err != nil {
		log.Fatal(err)

	}

	route := gin.Default()
	route.RedirectTrailingSlash = true

	route.GET("api_client/all", func(c *gin.Context) {
		var users []User
		var err error
		users, err = engine.GetAll(usersColl)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	route.Run(":3000")
}
