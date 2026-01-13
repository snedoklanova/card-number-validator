package main

import (
	"github.com/gin-gonic/gin"
	"github.com/snedoklanova/card-number-validator/service/api/album"
	"github.com/snedoklanova/card-number-validator/service/api/validator"
)

func main() {
	router := gin.Default()
  router.GET("/albums", album.GetAlbums)
	router.GET("/albums/:id", album.GetAlbumByID)
	router.GET("/validate/:number", validator.GetValidationByCardNumber)
	router.POST("/albums", album.PostAlbums)
	

  router.Run("localhost:8080")
}




