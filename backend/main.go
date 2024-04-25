package main

import (
	"github.com/danielargoe/programming-language-exploration-final/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/quotes", controllers.GetQuotes)
	r.GET("/quotes/:id", controllers.GetQuoteByID)

	r.PUT("/quotes/:id", controllers.UpdateQuote)
	r.POST("/quotes", controllers.CreateQuote)
	r.DELETE("/quotes/:id", controllers.DeleteQuote)

	r.Run()
}
