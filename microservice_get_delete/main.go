package main

import (
	"log"
	"net/http"
	user "microservice_get_delete/controller/users"

	"github.com/gin-gonic/gin"
)

//StartGin function
func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/users", user.GetAllUser)
		api.GET("/users/:id", user.GetUser)
		api.DELETE("/users/:id", user.DeleteUser)
	}
	router.NoRoute(func(c *gin.Context) {
		// c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusOK, gin.H{"status": "404", "msg": "Not Found"})
	})
	log.Println("[*] Starting Serve :8000 ...")
	log.Println("[+] Served ...")
	router.Run(":8000")
}
