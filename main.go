package main

import (
	configs "banking/config" // Asegúrate de que este sea el nombre correcto de tu paquete de configuración
	_ "banking/docs"         // Importa tu documentación de Swagger generada aquí
	"banking/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // Archivos estáticos para Swagger
	ginSwagger "github.com/swaggo/gin-swagger" // Gin-swagger para la documentación de la API
)

// @title API aplicacion B-Pagos
// @description Api para Banking.
// @version 1.0
// @host banking.nexus369.net
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Establece la conexión a la base de datos
	configs.ConnectToDB()

	r := gin.Default()

	

	// Configuración de CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(config))

	// Configura tus rutas usando la función SetupRouter de tu paquete routes
	routes.UserRouter(r)
	routes.BancoRouter(r)
	routes.TransaccionRouter(r) // Asumiendo que esta función acepta *gin.Engine y configura las rutas

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Inicia el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Puerto por defecto si no está definido en las variables de entorno
	}
	r.Run(":" + port)
}
