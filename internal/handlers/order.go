package handlers

import (
	"net/http"
	"project_go/internal/models"
	"project_go/internal/utils"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrder godoc
// @Summary Создать заказ
// @Description Добавляет заказ для указанного пользователя
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param input body models.OrderInput true "Данные заказа"
// @Success 201 {object} models.Order
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id}/orders [post]
// @Security BearerAuth
func CreateOrder(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := utils.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input models.OrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	order := models.Order{
		UserID:    uint(userID),
		Product:   input.Product,
		Quantity:  input.Quantity,
		Price:     input.Price,
		CreatedAt: time.Now(),
	}

	utils.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}

// GetOrders godoc
// @Summary Получить список заказов пользователя
// @Tags orders
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {array} models.Order
// @Router /users/{id}/orders [get]
// @Security BearerAuth
func GetOrders(c *gin.Context) {
	userID := c.Param("id")
	var orders []models.Order

	if err := utils.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
