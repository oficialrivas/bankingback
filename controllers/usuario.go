package controllers

import (
	"banking/config"
	"banking/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// CreateUser godoc
// @Summary Crea un nuevo usuario
// @Description Agrega un nuevo usuario a la base de datos
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param user body models.User true "Datos del Usuario"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{} "Datos de usuario no válidos"
// @Failure 500 {object} map[string]interface{} "Error al crear el usuario"
// @Router /usuarios [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario no válidos"})
		return
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsers godoc
// @Summary Lista todos los usuarios
// @Description Obtiene una lista de todos los usuarios registrados
// @Tags usuarios
// @Produce  json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{} "Error al obtener los usuarios"
// @Router /usuarios [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := configs.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary Obtiene un usuario por su ID
// @Description Obtiene detalles de un usuario específico por su ID único
// @Tags usuarios
// @Produce  json
// @Param id path string true "ID del Usuario"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{} "ID de usuario no válido"
// @Failure 404 {object} map[string]interface{} "Usuario no encontrado"
// @Router /usuarios/{id} [get]
func GetUserByID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}

	var user models.User
	if err := configs.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Actualiza un usuario por su ID
// @Description Actualiza los detalles de un usuario específico
// @Tags usuarios
// @Accept  json
// @Produce  json
// @Param id path string true "ID del Usuario"
// @Param user body models.User true "Datos del Usuario Actualizados"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{} "ID de usuario no válido"
// @Failure 400 {object} map[string]interface{} "Datos de usuario no válidos"
// @Failure 500 {object} map[string]interface{} "Error al actualizar el usuario"
// @Router /usuarios/{id} [put]
func UpdateUser(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}

	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario no válidos"})
		return
	}

	if err := configs.DB.Model(&models.User{}).Where("id = ?", userID).Updates(newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

// DeleteUser godoc
// @Summary Elimina un usuario por su ID
// @Description Elimina un usuario específico de la base de datos
// @Tags usuarios
// @Produce  json
// @Param id path string true "ID del Usuario"
// @Success 200 {object} map[string]interface{} "message: Usuario eliminado exitosamente"
// @Failure 400 {object} map[string]interface{} "ID de usuario no válido"
// @Failure 500 {object} map[string]interface{} "Error al eliminar el usuario"
// @Router /usuarios/{id} [delete]
func DeleteUser(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}

	if err := configs.DB.Where("id = ?", userID).Delete(&models.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}
