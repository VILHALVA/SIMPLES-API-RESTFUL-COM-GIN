package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "codigo/models"
)

var videos = []models.Video{
    {ID: 1, Title: "Inception", Director: "Christopher Nolan"},
    {ID: 2, Title: "The Matrix", Director: "The Wachowskis"},
}

// GetVideos retorna todos os vídeos
func GetVideos(c *gin.Context) {
    c.JSON(http.StatusOK, videos)
}

// GetVideoByID retorna um vídeo pelo ID
func GetVideoByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    for _, video := range videos {
        if video.ID == id {
            c.JSON(http.StatusOK, video)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
}

// CreateVideo cria um novo vídeo
func CreateVideo(c *gin.Context) {
    var newVideo models.Video
    if err := c.ShouldBindJSON(&newVideo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newVideo.ID = len(videos) + 1
    videos = append(videos, newVideo)
    c.JSON(http.StatusCreated, newVideo)
}

// UpdateVideo atualiza um vídeo existente
func UpdateVideo(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var updatedVideo models.Video
    if err := c.ShouldBindJSON(&updatedVideo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, video := range videos {
        if video.ID == id {
            videos[i].Title = updatedVideo.Title
            videos[i].Director = updatedVideo.Director
            c.JSON(http.StatusOK, videos[i])
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
}

// DeleteVideo deleta um vídeo
func DeleteVideo(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    for i, video := range videos {
        if video.ID == id {
            videos = append(videos[:i], videos[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Video deleted"})
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
}
