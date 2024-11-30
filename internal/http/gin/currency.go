package gin

import (
	"log"
	"net/http"

	"github.com/danielmalka/challenge-bravo/application/currency"
	"github.com/danielmalka/challenge-bravo/config"
	"github.com/danielmalka/challenge-bravo/pkg/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initService(config config.Config) (*gorm.DB, error, GinResponse, *currency.Service) {
	db, err, response := getDB(config)
	if err != nil {
		log.Println("error connecting to database: ", err)
		return nil, err, response, nil
	}
	service := currency.NewService(db)
	return db, err, response, service
}

func listCurrency(config config.Config) gin.HandlerFunc {
	db, err, response, service := initService(config)
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(response.StatusCode, response.Message)
		}
	}

	currencies, err := service.List()
	if err != nil {
		log.Println("error listing currencies: ", err)
		response.StatusCode = http.StatusInternalServerError
		response.Message = gin.H{"error": err.Error()}
		return func(c *gin.Context) {
			c.JSON(response.StatusCode, response.Message)
		}
	}

	storage.Close(db)
	return func(c *gin.Context) {
		response.StatusCode = http.StatusOK
		c.JSON(response.StatusCode, currencies)
	}
}

func createCurrency(config config.Config) gin.HandlerFunc {
	db, err, response, service := initService(config)
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(response.StatusCode, response.Message)
		}
	}

	return func(c *gin.Context) {
		var request currency.CreateData
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Println("error binding JSON: ", err)
			response.StatusCode = http.StatusBadRequest
			response.Message = gin.H{"error": err.Error()}
			c.JSON(response.StatusCode, response.Message)
			return
		}

		currency, err := service.Create(&request)
		if err != nil {
			log.Println("error creating currency: ", err)
			response.StatusCode = http.StatusInternalServerError
			response.Message = gin.H{"error": err.Error()}
			c.JSON(response.StatusCode, response.Message)
			return
		}

		storage.Close(db)
		response.StatusCode = http.StatusOK
		c.JSON(response.StatusCode, currency)
	}
}

func updateCurrency(config config.Config) gin.HandlerFunc {
	db, err, response, service := initService(config)
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(response.StatusCode, response.Message)
		}
	}

	return func(c *gin.Context) {
		var request currency.UpdateData
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Println("error binding JSON: ", err)
			response.StatusCode = http.StatusBadRequest
			// make a better response message
			response.Message = gin.H{"error": err.Error()}
			c.JSON(response.StatusCode, response.Message)
			return
		}
		request.ID = c.Param("id")

		currency, err := service.Update(&request)
		if err != nil {
			log.Println("error updating currency: ", err)
			response.StatusCode = http.StatusInternalServerError
			response.Message = gin.H{"error": err.Error()}
			c.JSON(response.StatusCode, response.Message)
			return
		}

		storage.Close(db)
		response.StatusCode = http.StatusOK
		c.JSON(response.StatusCode, currency)
	}
}

func deleteCurrency(config config.Config) gin.HandlerFunc {
	db, err, response, service := initService(config)
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(response.StatusCode, response.Message)
		}
	}

	return func(c *gin.Context) {
		currencyID := c.Param("id")

		err := service.Delete(currencyID)
		if err != nil {
			log.Println("error updating currency: ", err)
			response.StatusCode = http.StatusInternalServerError
			response.Message = gin.H{"error": err.Error()}
			c.JSON(response.StatusCode, response.Message)
			return
		}

		storage.Close(db)
		response.StatusCode = http.StatusOK
		c.JSON(response.StatusCode, gin.H{})
	}
}
