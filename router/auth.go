package router

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/leonardogcsoares/phonebook-api/config"
)

func (r Router) validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jwt authorization goes here
		tokenString := c.Request.Header.Get("Authorization")

		options, err := config.New()
		if err != nil {
			c.String(http.StatusInternalServerError, "invalid config/options: "+err.Error())
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return options.VerifyKey, nil
		})

		if err != nil {
			c.String(http.StatusUnauthorized, "unauthorized")
			return
		}
		if !token.Valid {
			c.String(http.StatusUnauthorized, "unauthorized")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			username := claims["username"].(string)
			password := claims["password"].(string)

			if username != "admin" && password != "admin" {
				c.String(http.StatusUnauthorized, "unauthorized")
				return
			}
		} else {
			c.String(http.StatusUnauthorized, "unauthorized")
			return
		}

		c.Next()
	}
}

func (r Router) login(c *gin.Context) {
	var lr loginReq
	err := c.BindJSON(&lr)
	if err != nil {
		c.String(http.StatusBadRequest, "login request not in valid format")
		return
	}

	options, err := config.New()
	if err != nil {
		c.String(http.StatusInternalServerError, "invalid config/options")
		return
	}

	// get ID for user to use as parser
	exp := time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp":      exp,
		"username": lr.Username,
		"password": lr.Password,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(options.SignKey)
	if err != nil {
		c.String(http.StatusInternalServerError, "unable to sign key")
		return
	}

	c.JSON(
		http.StatusOK,
		LoginResp{
			Token: tokenString,
		})

	return

}
