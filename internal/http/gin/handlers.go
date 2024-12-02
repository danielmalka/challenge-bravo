package gin

import (
	"net/http"

	"github.com/danielmalka/challenge-bravo/application/currency"
	"github.com/gin-gonic/gin"
)

func Handlers(appStage string, service *currency.Service, response GinResponse) *gin.Engine {
	if appStage == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(customLogger())
	r.Use(returnHeaders())

	r.GET("", getHome())

	v1 := r.Group("/v1")
	{
		// Currency Routes
		v1.GET("/currency", listCurrency(service, response))
		v1.POST("/currency", createCurrency(service, response))
		v1.PUT("/currency/:id", updateCurrency(service, response))
		v1.DELETE("/currency/:id", deleteCurrency(service, response))
		// Conversion Routes
		v1.GET("/conversion", doConversion(service, response))
	}
	return r
}

func getHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Challenge Bravo"})
	}
}
