package routes

import (
    "github.com/gin-gonic/gin"
    "banking/controllers"
)

func UserRouter(router *gin.Engine) {
    api := router.Group("/")
    {
        api.POST("/users", controllers.CreateUser)
        api.GET("/users", controllers.GetUsers)
        api.GET("/users/:id", controllers.GetUserByID)
        api.PUT("/users/:id", controllers.UpdateUser)
        api.DELETE("/users/:id", controllers.DeleteUser)
        api.GET("/users/query", controllers.GetUserByNameAndPin)
    }
}


	