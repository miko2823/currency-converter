package account

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/miko2823/currency-converter.git/config"
	"github.com/miko2823/currency-converter.git/domain/account/models"
	"github.com/miko2823/currency-converter.git/domain/account/repository"
)

type AccountUsecase interface {
	Login(user_name string, password string) (string, error)
	generateJWT(user *models.User) (string, error)
}

type accountUsecase struct {
	config            config.Environment
	accountRepository repository.AccountRepository
}

func NewAccountUsecase(config config.Environment, repo repository.AccountRepository) AccountUsecase {
	return accountUsecase{config, repo}
}

func (u accountUsecase) Login(user_name string, password string) (string, error) {
	user, err := u.accountRepository.Login(user_name, password)
	if err != nil {
		return "", err
	}
	token, err := u.generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
func (u accountUsecase) generateJWT(user *models.User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"name": user.UserName,
		"exp":  time.Now().Add(time.Duration(u.config.TOKEN_EXPIRATION) * time.Hour).Unix(),
	}).SignedString([]byte(u.config.TOKEN_SIGNING_KEY))
}
