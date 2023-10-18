package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "urban-map/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	r.GET("/markers/:id", v1.GetMarker)
	r.GET("/markers", v1.GetMarkers)
	return r
}