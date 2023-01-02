package converter

import (
	"database/sql"

	"github.com/miko2823/currency-converter.git/config"
	"github.com/miko2823/currency-converter.git/domain/converter/repository"
)

type converterPersistence struct {
	config config.Environment
	conn   *sql.DB
}

func NewConverterPersistence(config config.Environment, conn *sql.DB) repository.ConverterRepository {
	return converterPersistence{config, conn}
}
