package auth

import (
	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	jose "gopkg.in/square/go-jose.v2"
	"net/http"
)

func terminateWithError(statusCode int, message string, c *gin.Context) {
	c.JSON(statusCode, gin.H{"error": message})
	c.Abort()
}

func AuthRequired(domain string, audience string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Checking Auth")
		var auth0Domain = "https://" + domain + "/"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
		configuration := auth0.NewConfiguration(client, []string{audience}, auth0Domain, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)

		_, err := validator.ValidateRequest(c.Request)

		if err != nil {
			log.Println(err)
			terminateWithError(http.StatusUnauthorized, "token is not valid", c)
			return
		}

		c.Next()
	}
}
