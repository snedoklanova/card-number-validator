package main

import (
	"github.com/gin-gonic/gin"

	album "github.com/snedoklanova/credit-card-number-validator/service/api/album"
)

func main() {
	router := gin.Default()
  router.GET("/albums", album.GetAlbums)
	router.GET("/albums/:id", album.getAlbumByID)
	router.GET("/validate/:number", validator.getValidationByCardNumber)
	router.POST("/albums", album.postAlbums)
	

  router.Run("localhost:8080")
}




