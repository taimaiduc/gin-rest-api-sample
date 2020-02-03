package main

import (
	"os"

	"gin-rest-api-sample/api"
	"gin-rest-api-sample/database"
	"gin-rest-api-sample/lib/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
	gin.SetMode(os.Getenv("ENV"))
	// initializes database
	db, _ := database.Initialize()

	port := os.Getenv("PORT")
	app := gin.Default() // create gin app
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(app) // apply api router
	app.Run(":" + port)  // listen to given port
}
