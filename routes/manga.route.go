package routes

import (
	"manga-hub/handlers"

	"github.com/gin-gonic/gin"
)


func MangaRoutes(r *gin.RouterGroup) {

	r.POST("/download", handlers.MangaToPdf)
}