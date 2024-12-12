package main

import (
	"net/http"
	"strconv"

	"books-gin-api/model/services"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books := services.GetAllBooks()
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalidID"})
		return
	}
	book, err := services.GetBooksByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}
