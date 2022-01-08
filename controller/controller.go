package controller

import (
	"bookstore/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var books = []model.Books{
	{Title: "Harry Potter", Author: "J. K. Rowling"},
	{Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func GetBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func PostBook(c echo.Context) error {
	var book model.Books
	if err := c.Bind(&book); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	books = append(books, book)
	return c.JSON(http.StatusCreated, book)
}
