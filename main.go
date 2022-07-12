package main


import (
 "go-youtube/api"
 "go-youtube/config"
"github.com/gin-gonic/gin"
)

func main() {
// initialize new gin engine (for server)
 r := gin.Default()

// create/configure database instance
 db := config.CreateDatabase()

 // configure firebase
 firebaseAuth := config.SetupFirebase()

 r.Use(func(c *gin.Context) {
	c.Set("db", db)
	c.Set("firebaseAuth", firebaseAuth)
   })

// set db to gin context with a middleware to all incoming request
 r.Use(func(c *gin.Context) {
  c.Set("db", db)
 })

// routes definition for finding and creating artists
 r.GET("/artist", api.FindUsers)
 r.POST("/artist", api.CreateUser)

// start the server
 r.Run(":3000")
}