package main

import (
	"thelist/inits"
	"thelist/controllers"
	"thelist/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {

	r:= gin.Default()

	r.POST("/entries", middlewares.RequireAuth, controllers.CreateEntry)
	r.GET("/entries", controllers.GetEntries)
	r.GET("/entries/:entryId", controllers.GetEntry)
	r.PUT("/entries/:entryId", controllers.UpdateEntry)
	r.DELETE("/entries/:entryId", controllers.DeleteEntry)


	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/validate", controllers.Validate)
	r.GET("/users", controllers.GetUsers)
	r.POST("/logout", controllers.Logout)


	r.Run()
}