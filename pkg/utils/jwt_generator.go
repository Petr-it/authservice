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
	"golang.org/x/crypto/bcrypt"
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
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	encryptedUID, _ := encryptUID(userID, []byte(os.Getenv("UID_ENCRYPTION_KEY")))

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

func GenerateNewRefreshToken(userID string) (string, string, error) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return "", "", err
	}

	t := base64.URLEncoding.EncodeToString([]byte(b))

	hash, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return t, string(hash), nil
}
