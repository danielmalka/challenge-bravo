package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// TableName - force the table name
func (Currency) TableName() string {
	return "currency"
}

// Currency Table
type Currency struct {
	ID              uint   `gorm:"primary_key"`
	Code            string `gorm:"type:varchar(3);index"`
	Name            string `gorm:"type:varchar(60);"`
	BackingCurrency bool   `gorm:"default:false"`
	CurrencyRate    string `gorm:"type:varchar(20);"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

// repository struct
type repository struct {
	db *gorm.DB
}

// NewRepository creates a new currency repository
func NewRepository(db *gorm.DB) Repository {
	migrateAndSeed(db)
	return &repository{db: db}
}

// GetByID retrieves a currency by its ID
func (r *repository) GetByID(id *string) (*Currency, error) {
	var currency Currency
	if err := r.db.First(&currency, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &currency, nil
}

func (r *repository) List() (*ResponseList, error) {
	var currencies []Currency
	if err := r.db.Find(&currencies).Error; err != nil {
		return nil, err
	}

	currencyPointers := make([]*Currency, len(currencies))
	for i := range currencies {
		currencyPointers[i] = &currencies[i]
	}

	return &ResponseList{
		Currencies: currencyPointers,
	}, nil
}

func (r *repository) Create(code, name, currency_rate string, backing_currency bool) (*Currency, error) {
	newCurrency := Currency{
		Code:            code,
		Name:            name,
		BackingCurrency: backing_currency,
		CurrencyRate:    currency_rate,
	}
	result := r.db.Create(&newCurrency)
	return &newCurrency, result.Error
}

func (r *repository) Update(id, code, name, currency_rate string, backing_currency bool) (*Currency, error) {
	var existingCurrency Currency
	if err := r.db.First(&existingCurrency, "id = ?", &id).Error; err != nil {
		return nil, err
	}
	if notEmpty(code) {
		existingCurrency.Code = code
	}
	if notEmpty(name) {
		existingCurrency.Name = name
	}
	if notEmpty(currency_rate) {
		existingCurrency.CurrencyRate = currency_rate
	}
	existingCurrency.BackingCurrency = backing_currency
	result := r.db.Save(&existingCurrency)
	return &existingCurrency, result.Error
}

func (r *repository) Delete(id *string) error {
	return r.db.Delete(&Currency{}, "id = ?", id).Error
}

func (r *repository) GetByCodes(codes []string) (*ResponseList, error) {
	if len(codes) == 0 {
		return nil, errors.New("no codes provided")
	}

	var currencies []Currency
	if err := r.db.Where("code IN ?", codes).Find(&currencies).Error; err != nil {
		return nil, err
	}

	currencyPointers := make([]*Currency, len(currencies))
	for i := range currencies {
		currencyPointers[i] = &currencies[i]
	}

	return &ResponseList{
		Currencies: currencyPointers,
	}, nil
}

func migrateAndSeed(db *gorm.DB) {
	db.AutoMigrate(
		&Currency{},
	)
	// Add USD to currency table
	usd := Currency{
		Code:            "USD",
		Name:            "Dollar",
		CurrencyRate:    "1",
		BackingCurrency: true,
	}
	if db.Migrator().HasTable(&Currency{}) {
		if err := db.First(&usd).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&usd)
		}
	}
}
