package main

import (
    "github.com/gin-gonic/gin"
    "codigo/controllers"
)

func main() {
    r := gin.Default()

    // Rotas CRUD
    r.GET("/videos", controllers.GetVideos)
    r.GET("/videos/:id", controllers.GetVideoByID)
    r.POST("/videos", controllers.CreateVideo)
    r.PUT("/videos/:id", controllers.UpdateVideo)
    r.DELETE("/videos/:id", controllers.DeleteVideo)

    // Inicia o servidor na porta 8080
    r.Run(":8080")
}
