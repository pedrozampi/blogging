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

//PUT

func updatePost(c *gin.Context) {
	var updatedPost BlogPost
	id := c.GetInt("id")

	if err := c.BindJSON(&updatedPost); err != nil {
		return
	}

	for i := 0; i < len(PostList); i++ {
		if PostList[i].ID == id {
			PostList[i] = updatedPost
			c.JSON(http.StatusOK, updatedPost)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Post not found."})
}

// DELETE
func deletePost(c *gin.Context) {
	id := c.GetInt("id")

	for i := 0; i < len(PostList); i++ {
		if PostList[i].ID == id {
			PostList = append(PostList[:id], PostList[i+1:]...)
			c.JSON(http.StatusNoContent, gin.H{"message": "Post deleted."})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Post not found."})
}

func main() {
	router := gin.Default()

	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPostByID)

	router.POST("/posts", registerPost)

	router.PUT("posts/:id", updatePost)

	router.DELETE("posts/:id", deletePost)

	router.Run("localhost:8080")
}
