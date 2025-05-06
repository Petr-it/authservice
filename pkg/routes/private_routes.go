package routes

import (
	"authservice/app/controllers"
	"authservice/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/v1")

	route.Get("/uid", middleware.JWTProtected(), controllers.GetUid)
}
