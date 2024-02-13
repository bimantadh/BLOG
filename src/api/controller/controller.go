package controller

import (
	"fmt"
	"myproject/src/api/model"
	"myproject/src/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := model.Post{Title: input.Title, Content: input.Content}
	database.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func FindPost(c *gin.Context) {
	var post model.Post

	if err := database.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func UpdatePost(c *gin.Context) {
	var post model.Post
	if err := database.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var input CreatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	updatePost := model.Post{Title: input.Title, Content: input.Content}
	database.DB.Model(&post).Updates(&updatePost)
	c.JSON(http.StatusOK, gin.H{"data": post})

}

func DeletePost(c *gin.Context) {

	var post model.Post
	if err := database.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})	

	return
	}
	database.DB.Delete(&post)
	fmt.Println("id has been delete", post.ID)
	c.JSON(http.StatusOK, gin.H{"id has been delete": post.ID})
}
