package controller

import (
	"bookstore/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var books = []model.Books{
	{Id: 10001, Title: "Harry Potter", Author: "J. K. Rowling"},
	{Id: 10003, Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{Id: 10003, Title: "The Wizard of Oz", Author: "L. Frank Baum"},
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

func DeleteBook(c echo.Context) error {
	deleteID := c.Param("id")
	fmt.Println(deleteID)
	fmt.Println("-----------------")
	delId, err := strconv.ParseInt(deleteID, 10, 64)
	fmt.Println(delId)
	fmt.Println("-----------------")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for i, v := range books {
		if v.Id == delId {
			fmt.Println(v.Id)
			books = append(books[:i], books[i+1:]...)
			fmt.Println(books)
		}
	}
	return c.JSON(http.StatusNoContent, "Delected this contect")
}
