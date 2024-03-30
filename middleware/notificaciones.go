package middleware

import (

    "banking/models"
    "banking/utils"
    "github.com/gin-gonic/gin"
    "net/http"
)

func NotificarCambioDeSaldo() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next() // Procesa la solicitud primero.

        // Comprueba si la solicitud fue exitosa.
        if c.Writer.Status() == http.StatusOK {
            // Extrae la información de usuario y comercio de los valores pasados por el contexto de Gin.
            usuario, _ := c.Get("usuario")
            comercio, _ := c.Get("comercio")
            
            if usuarioDB, ok := usuario.(models.User); ok {
                mensajeUsuario := "Tu saldo ha sido actualizado."
                utils.EnviarNotificacion(usuarioDB, mensajeUsuario)
            }

            if comercioDB, ok := comercio.(models.User); ok { // Asumiendo que usas el mismo modelo para comercios.
                mensajeComercio := "Tu saldo ha sido actualizado debido a una transacción reciente."
                utils.EnviarNotificacion(comercioDB, mensajeComercio)
            }
        }
    }
}
