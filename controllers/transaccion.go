package controllers

import (
    
    "net/http"
    "github.com/gin-gonic/gin"
    "banking/config" // Asegúrate de que este es el nombre correcto de tu paquete de configuración
    "banking/models"
)



// ProcesarTransaccion godoc
// @Summary Procesa una transacción entre un usuario y un comercio
// @Description Debida un monto del saldo de un usuario y lo acredita a un comercio
// @Tags transacciones
// @Accept  json
// @Produce  json
// @Param   input  body  models.TransaccionInput  true  "Información de la Transacción"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /procesar-transaccion [post]
func ProcesarTransaccion(c *gin.Context) {
    var input models.TransaccionInput

    // Intenta vincular el JSON de entrada a la estructura TransaccionInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    usuarioCedula := input.Usuario.Cedula
    comercioRif := input.Comercio.Rif
    montoTransaccion := float32(input.Comercio.Monto)

    var usuarioDB, comercioDB models.User
    resultUser := configs.DB.Where("cedula = ?", usuarioCedula).First(&usuarioDB)
    resultComercio := configs.DB.Where("rif = ?", comercioRif).First(&comercioDB)

    // Verifica errores en la búsqueda
    if resultUser.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
        return
    }
    if resultComercio.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Comercio no encontrado"})
        return
    }

    // Verifica si el usuario tiene saldo suficiente
    if usuarioDB.Saldo < montoTransaccion {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Saldo insuficiente"})
        return
    }

   
    usuarioDB.Saldo -= montoTransaccion
  

  
    comercioDB.Saldo += montoTransaccion
   

    // Actualiza ambos registros dentro de una transacción para asegurar la atomicidad
    tx := configs.DB.Begin()
    if err := tx.Save(&usuarioDB).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo debitar el saldo del usuario"})
        return
    }

    if err := tx.Save(&comercioDB).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo acreditar el saldo al comercio"})
        return
    }

    // Confirma los cambios si todo sale bien
    tx.Commit()

    // Antes de responder, establece los datos actualizados en el contexto para el middleware
    c.Set("usuario", usuarioDB)
    c.Set("comercio", comercioDB)

    // Responde con la confirmación de la transacción
    c.JSON(http.StatusOK, gin.H{
        "message": "Transacción exitosa",
        "usuarioSaldoActual": usuarioDB.Saldo,
        "comercioSaldoActual": comercioDB.Saldo,
    })
}
