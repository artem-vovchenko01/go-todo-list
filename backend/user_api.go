package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var curUserId = 1

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
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": errWhileJWTCreate})
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
		SendCustomError(c, http.StatusBadRequest, errEmailOccupied)
		return
	}
	Storage.users.Add(user)
	SendResponse(c, 201, registeredMesg)
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

	newTkn, err := GetRefreshedJWT(claims)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": errWhileJWTRefresh})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"token": newTkn})
}
