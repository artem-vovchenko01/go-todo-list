package main

import (
	"net/http"
	"time"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const BEARER_SCHEMA = "Bearer "
const JWT_EXPIRATION_TIME = 1
var httpStatusMessages = map[int]string {
	http.StatusUnauthorized : "unauthorized",
	http.StatusBadRequest : "bad request",
}

var JwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tkn,_, err := ParseJWT(c)
		if err != nil || !tkn.Valid {
			SendError(c, http.StatusUnauthorized)
			c.Abort()
		}
	}
}

func ParseJWT(c *gin.Context) (*jwt.Token, *Claims, error) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			return nil, nil, errors.New("no authentication header")
		}
		if len(authHeader) <= len(BEARER_SCHEMA) {
			return nil, nil, errors.New("flawed Authorization header")
		}
		tknStr := authHeader[len(BEARER_SCHEMA):]
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		return tkn, claims, err
}

func CreateJWT(user User) (string, error) {
	expirationTime := time.Now().Add(JWT_EXPIRATION_TIME * time.Minute)
	
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

func GetRefreshedJWT(claims *Claims) (string, error) {
	expirationTime := time.Now().Add(JWT_EXPIRATION_TIME * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
