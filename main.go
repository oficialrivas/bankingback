package main

import (
	"banking/config" // Asegúrate de que este sea el nombre correcto de tu paquete de configuración
	_ "banking/docs" // Importa tu documentación de Swagger generada aquí
	"banking/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // Archivos estáticos para Swagger
	ginSwagger "github.com/swaggo/gin-swagger" // Gin-swagger para la documentación de la API
	"log"
	"os"
)

// @title API aplicacion B-Pagos
// @description Api para Banking.
// @version 1.0
// @host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Establece la conexión a la base de datos
	configs.ConnectToDB()

	r := gin.Default()

	// Establecer los proxies de confianza
	r.SetTrustedProxies([]string{"127.0.0.1"}) // Especifica aquí las direcciones IP de tus proxies confiables

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

	// Configura el endpoint para Swagger UI utilizando gin-swagger
    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Inicia el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Puerto por defecto si no está definido en las variables de entorno
	}
	r.Run(":" + port)
}
