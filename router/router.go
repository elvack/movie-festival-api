package router

import (
	"github.com/elvack/movie-festival-api/controller/admin"
	"github.com/elvack/movie-festival-api/controller/health"
	"github.com/elvack/movie-festival-api/controller/movie"
	"github.com/elvack/movie-festival-api/controller/root"
	"github.com/elvack/movie-festival-api/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(db database.DB) (err error) {
	adminController := admin.NewController(db.GormDb)
	healthController := health.NewController(db.SqlDb)
	movieController := movie.NewController(db.GormDb)
	rootController := root.NewController()
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	router.GET("", rootController.Index)
	adminsGroup := router.Group("admins")
	{
		authGroup := adminsGroup.Group("auth")
		{
			authGroup.POST("sign-in", adminController.SignIn)
			authGroup.DELETE("sign-out", authorize(db.GormDb), adminController.SignOut)
		}
		moviesGroup := adminsGroup.Group("movies")
		{
			moviesGroup.POST("", authorize(db.GormDb), movieController.Create)
		}
	}
	router.GET("health", healthController.Check)
	router.Static("public", "./public")
	usersGroup := router.Group("users")
	{
		moviesGroup := usersGroup.Group("movies")
		{
			moviesGroup.GET("", authorize(db.GormDb), movieController.List)
		}
	}
	return router.Run()
}
