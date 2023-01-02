package repository

import "github.com/miko2823/currency-converter.git/domain/account/models"

type AccountRepository interface {
	Login(user_name string, password string) (*models.User, error)
}
