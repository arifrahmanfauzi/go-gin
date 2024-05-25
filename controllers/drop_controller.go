package controllers

import (
	"context"
	"go-starter/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDrops(c *gin.Context) {
	db, _ := c.Get("db")
	client := db.(*mongo.Client)
	dropRepo := repositories.NewDropRepository(client)

	drops, err := dropRepo.GetAllDrops(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, drops)
}
