package comment

import (
	"net/http"
	"sort"

	"github.com/PraserX/fullstack-rookie/pkg/database"
	"github.com/PraserX/fullstack-rookie/pkg/database/model"
	"github.com/PraserX/fullstack-rookie/pkg/webserver/helpers"
	"github.com/gin-gonic/gin"
)

type Comment struct {
	Comment  string `json:"comment"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

func LocationComments(router *gin.RouterGroup, db *database.Database) {
	router.GET("/comments", func(context *gin.Context) {
		context.Header("Content-Type", "application/json")
		comments := db.GetComments()

		sort.SliceStable(comments, func(i, j int) bool {
			return comments[i].Timestamp.After(comments[j].Timestamp)
		})

		context.JSON(http.StatusOK, gin.H{"comments": comments})
	})
	router.POST("/comments", func(context *gin.Context) {
		var err error
		var user model.User
		var comment Comment

		context.Header("Content-Type", "application/json")

		if err = context.ShouldBindJSON(&comment); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "bad input"})
			return
		}

		if bad, _ := helpers.ContainsForbiddenChars(comment); bad {
			context.JSON(http.StatusBadRequest, gin.H{"error": "forbidden chars"})
		}

		if !db.UserExists(comment.Email) {
			db.AddUser(comment.Nickname, comment.Email)
		}

		if user, err = db.GetUser(comment.Email); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "somthing goes wrong :/"})
			return
		}

		db.AddComment(comment.Comment, user)
		context.JSON(http.StatusOK, gin.H{})
	})
}
