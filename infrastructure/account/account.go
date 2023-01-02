package account

import (
	"database/sql"

	"github.com/miko2823/currency-converter.git/config"
	"github.com/miko2823/currency-converter.git/domain/account/models"
	"github.com/miko2823/currency-converter.git/domain/account/repository"
)

type accountPersistence struct {
	config config.Environment
	conn   *sql.DB
}

func NewAccountPersistence(config config.Environment, conn *sql.DB) repository.AccountRepository {
	return accountPersistence{config, conn}
}

func (r accountPersistence) Login(user_name string, password string) (*models.User, error) {
	// account := models.Account{ID: "123", Name: "account1", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	// row := r.Conn.QueryRow("SELECT * FROM users WHERE id = %s", id)
	// return convertToAccount(row)
	user := models.User{Id: "123abc", UserName: "kaori", Email: "test@test.com"}
	return &user, nil
}

// func convertToAccount(row *sql.Row) (models.Account, error) {
// 	account := models.Account{ID: "123", Name: "Tokyo visit", CreatedAt: time.Now(), UpdatedAt: time.Now()}
// 	// err := row.Scan(&account.ID, &account.Name, &account.CreatedAt, &account.UpdatedAt)
// 	// if err != nil {
// 	// 	if err == sql.ErrNoRows {
// 	// 		return nil, nil
// 	// 	}
// 	// 	return nil, err
// 	// }
// 	return account, nil
// }
