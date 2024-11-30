package gin

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/danielmalka/challenge-bravo/application/conversion"
	"github.com/danielmalka/challenge-bravo/config"
	"github.com/danielmalka/challenge-bravo/pkg/storage"
	"github.com/gin-gonic/gin"
)

func doConversion(config config.Config) gin.HandlerFunc {
	db, err, response, currencyService := initService(config)
	if err != nil {
		return func(c *gin.Context) {
			c.JSON(response.StatusCode, response.Message)
		}
	}

	return func(c *gin.Context) {
		var request conversion.ConversionData
		fromQuery := c.Query("from")
		toQuery := c.Query("to")
		amountQuery := c.Query("amount")
		if err := validateQueryString(fromQuery, toQuery, amountQuery); err != nil {
			log.Println("error binding JSON: ", err)
			response.StatusCode = http.StatusBadRequest
			// make a better response message
			response.Message = gin.H{"error": err.Error()}
			c.JSON(response.StatusCode, response.Message)
			return
		}
		request.From = fromQuery
		request.To = toQuery
		request.Amount = amountQuery

		if request.From == request.To {
			response.StatusCode = http.StatusBadRequest
			response.Message = gin.H{"error": errors.New("from and to currencies must be different")}
			c.JSON(response.StatusCode, response.Message)
			return
		}

		currencies, err := currencyService.GetByCodes(request.From, request.To)
		if err != nil {
			log.Println("error getting currencies: ", err)
			response.StatusCode = http.StatusInternalServerError
			response.Message = gin.H{"error": err.Error()}
			c.JSON(response.StatusCode, response.Message)
		}
		storage.Close(db)

		conversionValues := conversion.ConversionValues{
			Amount: request.Amount,
		}
		for _, currency := range currencies {
			if currency.Code == request.From {
				conversionValues.From.Code = currency.Code
				conversionValues.From.CurrencyRate = currency.CurrencyRate
			}
			if currency.Code == request.To {
				conversionValues.To.Code = currency.Code
				conversionValues.To.CurrencyRate = currency.CurrencyRate
			}
		}
		conversionService := conversion.NewService()
		cconversionResponse, err := conversionService.ConvertMoney(&conversionValues)
		if err != nil {
			log.Println("error getting currencies: ", err)
			response.StatusCode = http.StatusInternalServerError
			response.Message = gin.H{"error": err.Error()}
			c.JSON(response.StatusCode, response.Message)
		}

		response.StatusCode = http.StatusOK
		c.JSON(response.StatusCode, cconversionResponse)
	}
}

func validateQueryString(from, to, amount string) error {
	if strings.TrimSpace(from) == "" {
		return errors.New("from are required")
	}
	if strings.TrimSpace(to) == "" {
		return errors.New("to are required")
	}
	if strings.TrimSpace(amount) == "" {
		return errors.New("amount are required")
	}
	return nil
}
