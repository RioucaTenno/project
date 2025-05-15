package handlers

import (
	"net/http"
	"project_go/internal/models"
	"project_go/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser godoc
// @Summary Создание пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param input body models.UserInput true "Информация о пользователе"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var input models.UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var exists models.User
	if err := utils.DB.Where("email = ?", input.Email).First(&exists).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:         input.Name,
		Email:        input.Email,
		Age:          input.Age,
		PasswordHash: string(hashed),
	}

	utils.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"age":   user.Age,
	})
}

// GetUsers godoc
// @Summary Получить список пользователей
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Номер страницы"
// @Param limit query int false "Количество на странице"
// @Param min_age query int false "Минимальный возраст"
// @Param max_age query int false "Максимальный возраст"
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
// @Security BearerAuth
func GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	minAge, _ := strconv.Atoi(c.DefaultQuery("min_age", "0"))
	maxAge, _ := strconv.Atoi(c.DefaultQuery("max_age", "200"))

	var users []models.User
	var total int64

	query := utils.DB.Model(&models.User{})
	if minAge > 0 {
		query = query.Where("age >= ?", minAge)
	}
	if maxAge < 200 {
		query = query.Where("age <= ?", maxAge)
	}
	query.Count(&total).
		Limit(limit).
		Offset((page - 1) * limit).
		Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"total": total,
		"users": users,
	})
}

// GetUserByID godoc
// @Summary Получить пользователя по ID
// @Tags users
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
// @Security BearerAuth
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Обновить пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param input body models.UpdateUserInput true "Обновлённые данные"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [put]
// @Security BearerAuth
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	utils.DB.Model(&user).Updates(models.User{
		Name:  input.Name,
		Email: input.Email,
		Age:   input.Age,
	})

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Удалить пользователя
// @Tags users
// @Param id path int true "ID пользователя"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /users/{id} [delete]
// @Security BearerAuth
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	utils.DB.Delete(&user)
	c.Status(http.StatusNoContent)
}
