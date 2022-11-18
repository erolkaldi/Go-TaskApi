package helpers

import (
	"log"
	"os"
	"time"

	"task-api/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secret_key string = os.Getenv("SECRET_KEY")

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(bytes)

}

func ComparePassword(password, hash string) (valid bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateTokens(email string, username string, usertype string, userid string) (access_token string, refresh_token string, err error) {
	claims := &models.TokenDetails{
		Email:    email,
		UserName: username,
		UserType: usertype,
		UserID:   userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 3).Unix(),
		},
	}
	refreshClaims := &models.TokenDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
		},
	}
	access_token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret_key))
	if err != nil {
		log.Panic(err.Error())
		return
	}
	refresh_token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret_key))
	if err != nil {
		log.Panic(err.Error())
		return
	}
	return access_token, refresh_token, nil

}
