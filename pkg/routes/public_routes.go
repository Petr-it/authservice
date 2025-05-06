package routes

import (
	"authservice/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/v1")

	route.Get("/token/new", controllers.GetNewAccessToken)
	route.Get("/refresh", controllers.RefreshToken)
}
