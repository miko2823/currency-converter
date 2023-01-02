package converter

import (
	"github.com/miko2823/currency-converter.git/domain/converter/models"
	"github.com/miko2823/currency-converter.git/domain/converter/repository"
)

type ConverterUsecase interface {
	GetLatestRates(base string, symbols string, amount int) (*models.LatestRates, error)
}

type converterUsecase struct {
	converterRepository repository.ConverterRepository
}

func NewConverterUsecase(repo repository.ConverterRepository) ConverterUsecase {
	return converterUsecase{repo}
}

func (u converterUsecase) GetLatestRates(base string, symbols string, amount int) (*models.LatestRates, error) {
	converter, err := u.converterRepository.GetLatestRates(base, symbols, amount)
	if err != nil {
		return nil, err
	}

	return converter, nil
}
