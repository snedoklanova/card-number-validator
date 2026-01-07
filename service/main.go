package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
  router.GET("/albums", album.getAlbums)
	router.GET("/albums/:id", album.getAlbumByID)
	router.GET("/validate/:number", validator.getValidationByCardNumber)
	router.POST("/albums", album.postAlbums)
	

  router.Run("localhost:8080")
}




