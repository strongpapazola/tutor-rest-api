package main

import (
	"log"
	"net/http"
	user "microservice_insert_update/controller/users"

	"github.com/gin-gonic/gin"
)

//StartGin function
func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/users", user.CreateUser)
		api.PUT("/users/:id", user.UpdateUser)
	}
	router.NoRoute(func(c *gin.Context) {
		// c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusOK, gin.H{"status": "404", "msg": "Not Found"})
	})
	log.Println("[*] Starting Serve :8000 ...")
	log.Println("[+] Served ...")
	router.Run(":8000")
}
