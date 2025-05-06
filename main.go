package main

import (
	"authservice/pkg/configs"
	"authservice/pkg/middleware"
	"authservice/pkg/routes"
	"authservice/pkg/utils"

	_ "authservice/docs"

	"github.com/gofiber/fiber/v2"
)

// @title						API
// @version					1.0
// @description				This is an auto-generated API Docs.
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.email				your@mail.com
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @BasePath					/authservice
func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)
	middleware.FiberMiddleware(app)

	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServer(app)
}
