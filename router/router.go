package router

import (
	"github.com/elvack/movie-festival-api/controller/admin"
	"github.com/elvack/movie-festival-api/controller/health"
	"github.com/elvack/movie-festival-api/controller/root"
	"github.com/elvack/movie-festival-api/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(db database.DB) (err error) {
	adminController := admin.NewController(db.GormDb)
	healthController := health.NewController(db.SqlDb)
	rootController := root.NewController()
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	router.GET("", rootController.Index)
	adminGroup := router.Group("admin")
	{
		authGroup := adminGroup.Group("auth")
		{
			authGroup.POST("sign-in", adminController.SignIn)
			authGroup.DELETE("sign-out", authorize(db.GormDb), adminController.SignOut)
		}
	}
	router.GET("health", healthController.Check)
	router.Static("public", "./public")
	return router.Run()
}
