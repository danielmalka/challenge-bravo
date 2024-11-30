package currency

import (
	"github.com/danielmalka/challenge-bravo/application/currency/repository"
	"gorm.io/gorm"
)

type Service struct {
	repo repository.Repository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		repo: repository.NewRepository(db),
	}
}

// GetCurrencyByID retrieves a currency by its ID
func (s *Service) Get(id string) (*Currency, error) {
	currency, err := s.repo.GetByID(&id)
	if err != nil {
		return nil, err
	}
	return dbToEntity(currency), nil
}

// GetCurrencyByID retrieves a currency by its ID
func (s *Service) List() ([]*Currency, error) {
	respList, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return dbToEntities(respList.Currencies), nil
}

// CreateCurrency creates a new currency
func (s *Service) Create(c *CreateData) (*Currency, error) {
	newCurrency, err := s.repo.Create(c.Code, c.Name, c.CurrencyRate, c.DecimalSeparatorN, c.BackingCurrency)
	return dbToEntity(newCurrency), err
}

// UpdateCurrency updates an existing currency
func (s *Service) Update(c *UpdateData) (*Currency, error) {
	existingCurrency, err := s.repo.Update(c.ID, c.Code, c.Name, c.CurrencyRate, c.DecimalSeparatorN, c.BackingCurrency)
	if err != nil {
		return nil, err
	}
	return dbToEntity(existingCurrency), err
}

// DeleteCurrency deletes a currency by its ID
func (s *Service) Delete(id string) error {
	return s.repo.Delete(&id)
}
