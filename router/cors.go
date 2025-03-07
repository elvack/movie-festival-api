package router

import "github.com/gin-contrib/cors"

var corsConfig = cors.Config{
	AllowAllOrigins: true,
	AllowHeaders: []string{},
	AllowMethods: []string{},
}
