package controllers

import (
	"banking/config" // Asegúrate de que este sea el nombre correcto de tu paquete de configuración
	"banking/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// CreateBanco godoc
// @Summary Crea un nuevo banco
// @Description Agrega un nuevo banco a la base de datos
// @Tags bancos
// @Accept json
// @Produce json
// @Param banco body models.Banco true "Datos del Banco"
// @Success 201 {object} models.Banco
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /bancos [post]
func CreateBanco(c *gin.Context) {
	var banco models.Banco
	if err := c.BindJSON(&banco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos del banco no válidos"})
		return
	}

	if err := configs.DB.Create(&banco).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el banco"})
		return
	}

	c.JSON(http.StatusCreated, banco)
}

// GetBancos godoc
// @Summary Lista todos los bancos
// @Description Obtiene una lista de todos los bancos registrados
// @Tags bancos
// @Produce json
// @Success 200 {array} models.Banco
// @Failure 500 {object} map[string]interface{}
// @Router /bancos [get]
func GetBancos(c *gin.Context) {
	var bancos []models.Banco
	if err := configs.DB.Find(&bancos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los bancos"})
		return
	}

	c.JSON(http.StatusOK, bancos)
}

// GetBancoByID godoc
// @Summary Obtiene un banco por su ID
// @Description Obtiene detalles de un banco específico por su ID único
// @Tags bancos
// @Produce json
// @Param id path string true "ID del Banco"
// @Success 200 {object} models.Banco
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /bancos/{id} [get]
func GetBancoByID(c *gin.Context) {
	bancoID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del banco no válido"})
		return
	}

	var banco models.Banco
	if err := configs.DB.Where("id = ?", bancoID).First(&banco).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banco no encontrado"})
		return
	}

	c.JSON(http.StatusOK, banco)
}

// UpdateBanco godoc
// @Summary Actualiza un banco por su ID
// @Description Actualiza los detalles de un banco específico
// @Tags bancos
// @Accept json
// @Produce json
// @Param id path string true "ID del Banco"
// @Param banco body models.Banco true "Datos del Banco Actualizados"
// @Success 200 {object} models.Banco
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /bancos/{id} [put]
func UpdateBanco(c *gin.Context) {
	bancoID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del banco no válido"})
		return
	}

	var nuevoBanco models.Banco
	if err := c.BindJSON(&nuevoBanco); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos del banco no válidos"})
		return
	}

	if err := configs.DB.Model(&models.Banco{}).Where("id = ?", bancoID).Updates(nuevoBanco).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el banco"})
		return
	}

	c.JSON(http.StatusOK, nuevoBanco)
}

// DeleteBanco godoc
// @Summary Elimina un banco por su ID
// @Description Elimina un banco específico de la base de datos
// @Tags bancos
// @Produce json
// @Param id path string true "ID del Banco"
// @Success 200 {object} map[string]interface{} "message: Banco eliminado exitosamente"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /bancos/{id} [delete]
func DeleteBanco(c *gin.Context) {
	bancoID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del banco no válido"})
		return
	}

	if err := configs.DB.Where("id = ?", bancoID).Delete(&models.Banco{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el banco"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Banco eliminado exitosamente"})
}
