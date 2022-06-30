package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
)

var toValidate = map[string]string{
	"aud": "api://default",
	"cid": os.Getenv("OKTA_CLIENT_ID"),
}

func verify(c *gin.Context) bool {
	status := true
	token := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
		verifierSetup := jwtverifier.JwtVerifier{
			Issuer:           "https://" + os.Getenv("OKTA_DOMAIN") + "/oauth2/default",
			ClaimsToValidate: toValidate,
		}
		verifier := verifierSetup.New()
		_, err := verifier.VerifyAccessToken(token)
		if err != nil {
			c.String(http.StatusForbidden, err.Error())
			print(err.Error())
			status = false
		}
	} else {
		c.String(http.StatusUnauthorized, "Unauthorized")
		status = false
	}
	return status
}
func getUserData(c *gin.Context) {
	customerId := c.Param("customerId")
	fmt.Println(customerId)

	customer := Customer{
		ResponseData: ResponseDataObj{
			Id:            660903456,
			Email:         "murthy@svanpro.com",
			FirstName:     "Murthy",
			LastName:      "Avanithsaa",
			Country:       "IN",
			RegDate:       "2022-06-22 17:44:21",
			LastLoginDate: "2022-06-29 13:14:10",
			LastUserIp:    "49.37.151.250",
			ExternalId:    "",
			ExternalData: ExternalDataObj{
				History:   []HistoryObj{{"movie, Comedy", "Big Buck Bunny", "awWEFyPu", 596.5, 0.1676646974015088}},
				Favorites: []string{},
			},
		},
		Errors: []string{},
	}
	c.IndentedJSON(http.StatusOK, customer)
}

func main() {
	router := gin.Default()
	router.GET("/customers/:customerId", getUserData)

	router.Run("localhost:8080")
}

// Tags:     "movie,Comedy",
// Title:    "Big Buck Bunny",
// Mediaid:  "awWEFyPu",
// Duration: 596.5,
// Progress: 0.1676646974015088,
