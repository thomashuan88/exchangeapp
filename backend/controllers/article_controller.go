package controllers

import (
	"exchangeapp/backend/global"
	"exchangeapp/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArticle(ctx *gin.Context) {
	var Articale models.Articale

	if err := ctx.ShouldBindJSON(&Articale); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(&Articale); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.Create(&Articale).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": Articale})
}

func GetArticle(ctx *gin.Context) {
	var articales []models.Articale
	if err := global.Db.Find(&articales).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": articales})
}
