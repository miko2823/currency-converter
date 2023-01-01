package repository

import "github.com/miko2823/currency-converter.git/domain/converter/models"

type ConverterRepository interface {
	GetLatestRates(base string, symbols string, amount int) (models.LatestRates, error)
}
