package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserName     string `json:"user_name"`
	UserType     string `json:"user_type"`
	UserID       string `json:"user_id"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenDetails struct {
	UserName string `json:"user_name"`
	UserType string `json:"user_type"`
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
