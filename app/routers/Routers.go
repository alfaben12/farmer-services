package Routers

import (
	"net/http"
	"../Controllers/FarmLandController"
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token_result := strings.Split(tokenString, " ")[1]
	
	
	token, err := jwt.Parse(token_result, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		
		return []byte("ZPay"), nil
	})
	
	if token != nil && err == nil {
		fmt.Println("token verified")
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
		}
		
		return r
	}
	