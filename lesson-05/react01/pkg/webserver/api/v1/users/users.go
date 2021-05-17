package user

import (
	"net/http"

	"github.com/PraserX/fullstack-rookie/pkg/database"
	"github.com/gin-gonic/gin"
)

func LocationUsers(router *gin.RouterGroup, db *database.Database) {
	router.GET("/users", func(context *gin.Context) {
		context.Header("Content-Type", "application/json")
		users := db.GetUsers()
		context.JSON(http.StatusOK, gin.H{"users": users})
	})
}
