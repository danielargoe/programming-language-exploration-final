package controllers

import (
	"net/http"
	"strconv"

	"github.com/danielargoe/programming-language-exploration-final/models"
	"github.com/danielargoe/programming-language-exploration-final/repository"
	"github.com/gin-gonic/gin"
)

func GetQuotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repository.Quotes)
}

func GetQuoteByID(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(repository.Quotes); i++ {
		id, err := strconv.Atoi(id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
			return
		}

		if repository.Quotes[i].ID == id {
			c.IndentedJSON(http.StatusOK, repository.Quotes[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Quote not found"})
}

func CreateQuote(c *gin.Context) {
	var q models.Quote
	if err := c.BindJSON(&q); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}

	repository.Quotes = append(repository.Quotes, q)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Quote successfully added"})
}

func DeleteQuote(c *gin.Context) {
	id := c.Param("id")

	var index int
	for i := 0; i < len(repository.Quotes); i++ {
		id, err := strconv.Atoi(id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
			return
		}

		if repository.Quotes[i].ID == id {
			index = i
			break
		}
	}

	repository.Quotes = append(repository.Quotes[:index], repository.Quotes[index+1:]...)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Quote successfully removed"})
}

func UpdateQuote(c *gin.Context) {
	var q models.Quote
	if err := c.BindJSON(&q); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}

	id := c.Param("id")

	for i := 0; i < len(repository.Quotes); i++ {
		id, err := strconv.Atoi(id)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
			return
		}

		if repository.Quotes[i].ID == id {
			// repository.Quotes[i].ID = q.ID
			repository.Quotes[i].Author = q.Author
			repository.Quotes[i].Quote = q.Quote

			break
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Quote successfully updated"})
}
