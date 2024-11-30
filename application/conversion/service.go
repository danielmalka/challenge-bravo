package conversion

import (
	"fmt"

	"github.com/Rhymond/go-money"
	"github.com/shopspring/decimal"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// ConvertMoney converts money from one currency to another
func (s *Service) ConvertMoney(c *ConversionValues) (*ConversionResponse, error) {
	amount, err := decimal.NewFromString(c.Amount)
	if err != nil {
		return nil, err
	}
	if c.From.CurrencyRate == c.To.CurrencyRate {
		return getConversionResponse(amount, c)
	}

	// make better validation from backing currency getting from database
	if c.From.Code == money.USD {
		amountTo := getAmountFromRate(c.To.CurrencyRate, amount)
		return getConversionResponse(amountTo, c)
	}

	
	amountFromToUSD := getAmountFromRate(c.From.CurrencyRate, amount)
	amountTo := getAmountFromRate(c.To.CurrencyRate, amountFromToUSD)
	return getConversionResponse(amountTo, c)
}

func getAmountFromRate(currentRate string, amount decimal.Decimal) decimal.Decimal {
	currencyRate, _ := decimal.NewFromString(currentRate)
	return amount.Mul(currencyRate)
}

func convertAmountToInt(amount decimal.Decimal) int64 {
	digits := amount.NumDigits()
	factor := 10 * digits
	amountMul := amount.Mul(decimal.NewFromInt(int64(factor)))
	return amountMul.IntPart()
}

func getConversionResponse(amount decimal.Decimal, c *ConversionValues) (*ConversionResponse, error) {
	amountInt := convertAmountToInt(amount)
	formattedAmount := money.New(amountInt, c.To.Code)
	return &ConversionResponse{
		Description: fmt.Sprintf("Conversion from %s to %s", c.From.Code, c.To.Code),
		Amount:      amount.String(),
		Formatted:   formattedAmount.Display(),
	}, nil
}
