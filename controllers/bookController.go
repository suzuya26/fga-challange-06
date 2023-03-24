package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	err := ctx.ShouldBindJSON(&newBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newBook.ID = fmt.Sprintf("c%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, v := range BookDatas {
		if bookID == v.ID {
			condition = true
			BookDatas[i] = updatedBook
			BookDatas[i].ID = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully updated", bookID),
	})

}

func GetAllBook(ctx *gin.Context) {
	if len(BookDatas) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": "book haven't added yet",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": BookDatas,
	})
}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false
	var result Book

	for i, v := range BookDatas {
		if bookID == v.ID {
			condition = true
			result = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": result,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false
	var bookIndex int

	for i, v := range BookDatas {
		if bookID == v.ID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully deleted", bookID),
	})
}
