package controllers

import (
	"authservice/app/models"
	"authservice/pkg/utils"
	"authservice/platform/database"
	"time"

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
	userID := c.Query("uid")
	accessToken, err := utils.GenerateNewAccessToken(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "failed to generate access token: " + err.Error(),
		})
	}

	refreshToken, hash, err := utils.GenerateNewRefreshToken(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "failed to generate refresh token: " + err.Error(),
		})
	}

	reftoken := &models.RefreshToken{}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()
	reftoken.UserID = userID
	reftoken.TokenHash = hash
	reftoken.CreatedAt = time.Now()

	if err := validate.Struct(reftoken); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	if err := db.CreateToken(reftoken); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":         false,
		"msg":           nil,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// RefreshToken method for refresh a new access token.
//
//	@Description	refresh access token
//	@Summary		refresh access token
//	@Tags			Token
//	@Accept			json
//	@Produce		json
//	@Param			accesstoken	query		string	true	"User ID to encode into token"
//	@Param			refreshtoken	query		string	true	"User ID to encode into token"
//
// @Security ApiKeyAuth
//
//	@Success		200	{string}	status	"ok"
//	@Router			/v1/refresh [get]
func RefreshToken(c *fiber.Ctx) error {
	// accesstoken := c.Query("accesstoken")
	// refreshtoken := c.Query("refreshtoken")

	// claims, err := utils.VerifyAccessToken(accesstoken)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "Invalid access token",
	// 	})
	// }

	// userID := claims.UserID

	// db, err := database.OpenDBConnection()
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "DB connection failed: " + err.Error(),
	// 	})
	// }

	// storedToken, err := db.GetRefreshTokenByUserID(userID)
	// if err != nil || storedToken == nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "Refresh token not found",
	// 	})
	// }

	// if !utils.CompareHashWithToken(storedToken.TokenHash, refreshtoken) {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "Refresh token mismatch",
	// 	})
	// }

	// newAccessToken, err := utils.GenerateNewAccessToken(userID)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "Failed to generate access token",
	// 	})
	// }

	// newRefreshToken, newHash, err := utils.GenerateNewRefreshToken(userID)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "Failed to generate refresh token",
	// 	})
	// }

	// if err := db.UpdateRefreshToken(userID, newHash); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "Failed to update refresh token",
	// 	})
	// }

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		// "access_token":  newAccessToken,
		// "refresh_token": newRefreshToken,
	})
}

// GetUid method for get UID.
//
//	@Description	get UID
//	@Summary		get UID
//	@Tags			UID
//	@Accept			json
//	@Produce		json
//	@Param			refreshtoken	query		string	true	"User ID to encode into token"
//
// @Security ApiKeyAuth
//
//	@Success		200	{string}	status	"ok"
//	@Router			/v1/uid [get]
func GetUid(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		// "uid":   uid,
	})
}
