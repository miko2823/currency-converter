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

// func (r converterPersistence) Get(id string) (models.Converter, error) {
// 	// TODO connenct repo
// 	// converter := models.Converter{ID: "123", Name: "converter1", CreatedAt: time.Now(), UpdatedAt: time.Now()}
// 	// row := r.Conn.QueryRow("SELECT * FROM users WHERE id = %s", id)
// 	// return convertToConverter(row)
// 	return models.Converter{ID: "123", Name: "Tokyo visit", CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil
// }

// func convertToConverter(row *sql.Row) (models.Converter, error) {
// 	converter := models.Converter{ID: "123", Name: "Tokyo visit", CreatedAt: time.Now(), UpdatedAt: time.Now()}
// 	// err := row.Scan(&converter.ID, &converter.Name, &converter.CreatedAt, &converter.UpdatedAt)
// 	// if err != nil {
// 	// 	if err == sql.ErrNoRows {
// 	// 		return nil, nil
// 	// 	}
// 	// 	return nil, err
// 	// }
// 	return converter, nil
// }
