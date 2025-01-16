package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GET
func getPosts(c *gin.Context) {
	if c.GetString("term") != "" {
		getPostByTerm(c)
		return
	}
	c.JSON(http.StatusOK, PostList)
}

func getPostByID(c *gin.Context) {
	id := c.GetInt("id")

	for _, p := range PostList {
		if p.ID == id {
			c.JSON(http.StatusFound, p)
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Post not found."})
}

func getPostByTerm(c *gin.Context) {
	term := c.GetString("term")

	for _, p := range PostList {
		content := strings.Split(p.Content, " ")
		for _, con := range content {
			if con == term {
				c.JSON(http.StatusFound, p)
			}
		}

		if p.Category == term {
			c.JSON(http.StatusFound, p)
		}

		for _, t := range p.Tags {
			if t == term {
				c.JSON(http.StatusFound, p)
			}
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Post not found."})
}

// POST
func registerPost(c *gin.Context) {
	var newPost BlogPost

	if err := c.BindJSON(&newPost); err != nil {
		return
	}

	newPost.ID = ai.ID()

	PostList = append(PostList, newPost)
	c.JSON(http.StatusCreated, newPost)
}

func main() {
	router := gin.Default()

	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPostByID)

	router.POST("/posts", registerPost)

	router.Run("localhost:8080")
}
