package routes

import (
    "github.com/gin-gonic/gin"
    "banking/controllers"
)

func BancoRouter(router *gin.Engine) {
    

	api := router.Group("/")
    {
        // Rutas para operaciones CRUD de bancos
        api.POST("/bancos", controllers.CreateBanco)
        api.GET("/bancos", controllers.GetBancos)
        api.GET("/bancos/:id", controllers.GetBancoByID)
        api.PUT("/bancos/:id", controllers.UpdateBanco)
        api.DELETE("/bancos/:id", controllers.DeleteBanco)
    }
}