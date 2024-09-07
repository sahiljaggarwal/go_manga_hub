package main

import (
	"log"
	"manga-hub/config"
	"manga-hub/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	// "config/config.go"
)

func main(){
	gin.SetMode(gin.ReleaseMode)
	config.LoadConfig()

	router := gin.Default()
	port := config.Port
	if port == "" {
		port  = "3002"
	}
	server := &http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Testing route",
			"port":port,
		})
	})

	api := router.Group("/api/v1")
	routes.MangaRoutes(api)

	log.Println("Starting server on port:", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Error on starting server")
	}
	log.Println("Server is running on port: ", config.Port)



}