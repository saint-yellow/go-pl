package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/saint-yellow/go-pl/toys/memorandum/configuration"
)

var jWTSecret = []byte(configuration.JWT().Secret)

type Claims struct {
	jwt.StandardClaims

	UserID uint `json:"userID"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Authority int `json:"authority"`
}

func GenerateJWT(userID uint, userName, password string, authority int) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)
	claims := Claims{
		UserID: userID,
		UserName: userName,
		Password: password,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "memorandum",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jWTSecret)
}

func ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, 
		&Claims{}, 
		func(token *jwt.Token) (interface{}, error) {
			return jWTSecret, nil
		},
	)
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func ValidateJWT() gin.HandlerFunc {
	return func (c *gin.Context) {
		var code int
		var data interface{}

		token := c.GetHeader("Authorization")
		if token == "" {
			code = 1
		} else {
			claims, err := ParseJWT(token)
			if err != nil {
				code = 2
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 3
			}
		}

		if code != 0 {
			c.JSON(200, gin.H{
				"code": code,
				"message": "failed to validate JWT",
				"data": data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}