package Routers

import (
	"fmt"
	"net/http"
	"strings"

	"../Controllers/FarmLandController"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token_result := strings.Split(tokenString, " ")[1]

	token, err := jwt.Parse(token_result, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("ZPay"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Set("accountid", claims["accountid"])
		// c.JSON(http.StatusUnauthorized, claims["accountid"])
		c.Next()
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("farmlands", auth, FarmLandController.ListLand)
		v1.POST("farmlands", auth, FarmLandController.Create)
	}

	return r
}
