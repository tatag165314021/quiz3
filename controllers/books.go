package controllers

import (
	"net/http"
	"net/url"
	"quiz3/database"
	"quiz3/repository"
	"quiz3/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetAllBooks(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBooks(c *gin.Context) {
	var books structs.Books

	err := c.ShouldBindJSON(&books)
	if err != nil {
		panic(err)
	}

	cekurl, err2 := url.ParseRequestURI(books.ImageURL)
	if err2 != nil && books.ReleaseYear < 1980 && books.ReleaseYear > 2021 {
		panic("error url dan data diluar jangkauan")
	} else if err2 != nil {
		panic(err2)
	} else if books.ReleaseYear < 1980 && books.ReleaseYear > 2021 {
		panic("data diluar jangkauan")
	}

	books.ImageURL = cekurl.String()

	if books.TotalPage <= 100 {
		books.Thickness = "tipis"
	} else if books.TotalPage >= 101 && books.TotalPage <= 200 {
		books.Thickness = "sedang"
	} else if books.TotalPage >= 201 {
		books.Thickness = "tebal"
	}

	err = repository.InsertBooks(database.DbConnection, books)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Books" + books.Thickness,
	})
}

func UpdateBooks(c *gin.Context) {
	var books structs.Books
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&books)
	if err != nil {
		panic(err)
	}

	cekurl, err2 := url.ParseRequestURI(books.ImageURL)
	if err2 != nil && books.ReleaseYear < 1980 && books.ReleaseYear > 2021 {
		panic("error url dan data diluar jangkauan")
	} else if err2 != nil {
		panic(err2)
	} else if books.ReleaseYear < 1980 && books.ReleaseYear > 2021 {
		panic("data diluar jangkauan")
	}

	books.ImageURL = cekurl.String()

	if books.TotalPage <= 100 {
		books.Thickness = "tipis"
	} else if books.TotalPage >= 101 && books.TotalPage <= 200 {
		books.Thickness = "sedang"
	} else if books.TotalPage >= 201 {
		books.Thickness = "tebal"
	}

	books.ID = int64(id)
	err = repository.UpdateBooks(database.DbConnection, books)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Books",
	})
}

func DeleteBooks(c *gin.Context) {
	var books structs.Books
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	books.ID = int64(id)

	err = repository.DeleteBooks(database.DbConnection, books)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success delete Books",
	})

}

func GetBooksCategory(c *gin.Context) {
	var book structs.Books
	var books []structs.Books
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	book.CategoryID = int(id)

	var (
		result gin.H
	)

	books, err = repository.GetBooksCategory(database.DbConnection, book)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}
