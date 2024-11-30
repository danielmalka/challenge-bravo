package conversion

type ConversionData struct {
	From   string `json:"from" binding:"required"`
	To     string `json:"to" binding:"required"`
	Amount string `json:"amount" binding:"required"`
}

type ConversionValues struct {
	From   CurrencyRate
	To     CurrencyRate
	Amount string
}

type CurrencyRate struct {
	Code         string `json:"code"`
	CurrencyRate string `json:"currency_rate"`
}

type ConversionResponse struct {
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Formatted   string `json:"formatted"`
}

type UseCase interface {
	ConvertMoney(c *ConversionValues) (*ConversionResponse, error)
}
