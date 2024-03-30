package routes

import (
	"banking/controllers"
	"banking/middleware"
	"github.com/gin-gonic/gin"
)

func TransaccionRouter(router *gin.Engine) {
	router.POST("/procesar-transaccion", middleware.NotificarCambioDeSaldo(), controllers.ProcesarTransaccion)
}
