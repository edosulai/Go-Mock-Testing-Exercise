package controllers

import (
	"chal8/database"
	"chal8/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(product)

	product.UserID = uint(userData["id"].(float64))

	err = db.Create(&product).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if strings.ToLower(userData["role"].(string)) != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to perform this action",
		})
		return
	}
	product.UserID = uint(userData["id"].(float64))

	err = db.Model(&product).Where("id=?", productID).Updates(models.Product{Title: product.Title, Description: product.Description}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
func GetProduct(ctx *gin.Context) {
	db := database.GetDB()
	product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	err := db.Preload("User").First(&product, productID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Product not found",
		})
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func DeleteProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	err := db.Delete(&models.Product{}, productID).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if strings.ToLower(userData["role"].(string)) != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to perform this action",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
