package main

import (
	"log"
	"Book_talent/engine"
	//"github.com/PabloOsorix/Book_Talent/engine"

	"net/http"


	"github.com/gin-gonic/gin"
)

func main() {
	type User = engine.User
	client, err := engine.Create()
	if err != nil {
		log.Fatal(err)
	}
	usersColl, err := engine.Collection(client)
	if err != nil {
		log.Fatal(err)
	}

	app := gin.Default()
	app.RedirectTrailingSlash = true

	app.GET("/api/bot/all", func(c *gin.Context) {
		var users []User
		var err error
		users, err = engine.GetAll(usersColl)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})
	/*
	 End point to delete a record in the dabase, return 400 or 500 in case of fail.
	*/
	app.DELETE("/api/bot", func(c *gin.Context) {
		link := c.Query("link")
		if link == "" {
			c.JSON(400, gin.H{"error": "400", "message": "Missing parameter Link"})
			return
		}
		result, err := engine.Delete(usersColl, string(link))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	app.POST("/api/bot", func(c *gin.Context) {
		var new_user User
		new_user.Init()
		if err := c.ShouldBindJSON(&new_user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if new_user.Link == "" {
			c.JSON(400, gin.H{"error": "400", "message": "Missing parameter Link, Link is mandatory"})
			return
		} else if new_user.Name == "" {
			c.JSON(400, gin.H{"error": "400", "message": "Missing parameter Name, Name is mandatory"})
			return
		}
		response, err := engine.New(usersColl, new_user)
		if err != nil {
			c.JSON(500, err)
		}
		c.JSON(http.StatusOK, "New user "+response+" successfully created")
	})

	app.PUT("/api/bot", func(c *gin.Context) {
		var user_update User
		if err := c.ShouldBindJSON(&user_update); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else if user_update.Link == "" {
			c.JSON(400, gin.H{"error": "400", "message": "Missing parameter Link"})
			return
		} else if user_update.Name == "" {
			c.JSON(http.StatusNoContent, gin.H{"Error": "400", "message": "Missing parameter Name"})
			return
		} else if user_update.Experience == nil {
			c.JSON(400, gin.H{"error": "400", "message": "Missing Parameter Experience"})
			return
		}
		result, err := engine.Update(usersColl, user_update.Link, user_update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, result)
	})

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"400": "PAGE_NOT_FOUND", "error": "Page not found"})
	})

	defer engine.Disconnect(client)

	app.Run("localhost:5000")
}
