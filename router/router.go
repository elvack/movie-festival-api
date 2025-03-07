package router

import (
	"github.com/elvack/movie-festival-api/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(db database.DB) (err error) {
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	router.Static("public", "./public")
	return router.Run()
}
