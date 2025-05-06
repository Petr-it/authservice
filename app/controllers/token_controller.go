package controllers

import (
	"authservice/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// GetNewAccessToken method for create a new access token.
//
//	@Description	Create a new access token.
//	@Summary		create a new access token
//	@Tags			Token
//	@Accept			json
//	@Produce		json
//	@Param			uid	query		string	true	"User ID to encode into token"
//	@Success		200	{string}	status	"ok"
//	@Router			/v1/token/new [get]
func GetNewAccessToken(c *fiber.Ctx) error {
	return handleTokenGeneration(c)
}

// RefreshToken method for refresh a new access token.
//
//	@Description	refresh access token
//	@Summary		refresh access token
//	@Tags			Token
//	@Accept			json
//	@Produce		json
//
//	@Security		ApiKeyAuth
//
//	@Success		200	{string}	status	"ok"
//	@Router			/v1/refresh [get]
func RefreshToken(c *fiber.Ctx) error {
	return handleTokenGeneration(c)
}

func handleTokenGeneration(c *fiber.Ctx) error {
	userID := c.Query("uid")
	token, err := utils.GenerateNewAccessToken(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}
