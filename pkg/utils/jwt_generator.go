package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func encryptUID(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func GenerateNewAccessToken(userID string) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	encryptionKey := []byte(os.Getenv("UID_ENCRYPTION_KEY"))

	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	encryptedUID, _ := encryptUID(userID, encryptionKey)

	claims := jwt.MapClaims{
		"jti": encryptedUID,
		"exp": time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
