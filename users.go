package main

import (
	"net/http"
	"time"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var curUserId = 1

// var Users = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
// }

func Signin(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return
	}

	userId, err := Storage.userIds.Get(user.Email)
	var expectedPassword string
	if err == nil {
		existingUser, err := Storage.users.Get(userId)
		fmt.Println(existingUser)
		if err == nil {
			expectedPassword = existingUser.Password
		}
	}
	fmt.Println(user)

	if err != nil || expectedPassword != user.Password {
		SendError(c, http.StatusUnauthorized)
		return
	}

	tokenString, err := CreateJWT(user)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H {"message" : "error while creating JWT"})
		return
	}

	SendResponse(c, 200, tokenString)
}

func Signup(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return
	}

	if _, err := Storage.userIds.Get(user.Email); err == nil {
		SendCustomError(c, http.StatusBadRequest, "user with such email already exists")
			return
	}
		Storage.users.Add(user)
		SendResponse(c, 201, "registered")
	}

func Refresh(c *gin.Context) {
	tkn, claims, err := ParseJWT(c)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			SendError(c, http.StatusUnauthorized)
			return
		}
		SendError(c, http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		SendError(c, http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		SendError(c, http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	newTkn, err := GetRefreshedJWT(claims)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H {"message" : "internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H {"token" : newTkn })
}
