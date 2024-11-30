package gin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danielmalka/challenge-bravo/config"
	"github.com/danielmalka/challenge-bravo/pkg/healthcheck"
	"github.com/gin-gonic/gin"
)

func Handlers(config config.Config) *gin.Engine {
	if config.AppStage == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(customLogger())
	r.Use(returnHeaders())

	r.GET("/", getHome())
	r.GET("/health", healthCheck(config))

	v1 := r.Group("/v1")
	{
		cg := v1.Group("/currency")
		{
			cg.GET("/", listCurrency(config))
			cg.POST("/", createCurrency(config))
			cg.PUT("/:id", updateCurrency(config))
			cg.DELETE("/:id", deleteCurrency(config))
		}
		cg := v1.Group("/conversion")
		{
			cg.GET("/", doConversion(config))
		}
	}
	return r
}

func getHome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Challenge Bravo"})
	}
}

func healthCheck(config config.Config) gin.HandlerFunc {
	log.Printf("Health checking... ")

	statusCode := http.StatusOK
	msg := "✔️ Passed"
	userPass := fmt.Sprintf("%s:%s", config.DBUser, config.DBPass)
	host := fmt.Sprintf("%s:%s", config.DBHost, config.DBPort)
	err := healthcheck.HealthCheck(userPass, config.DBSchema, host)
	if err != nil {
		statusCode = http.StatusServiceUnavailable
		msg = fmt.Sprintf("❌ Failed with error: %s", err)
	}
	return func(c *gin.Context) {
		c.JSON(statusCode, gin.H{"message": msg})
	}
}
