package account

import (
	"fmt"
	"net/http"

	"github.com/miko2823/currency-converter.git/pkg"
	usecase "github.com/miko2823/currency-converter.git/usecase/account"
)

type AccountHandler interface {
	RegisterHandlers() http.Handler
	Login(w http.ResponseWriter, r *http.Request)
}

type accountHandler struct {
	accountUsecase usecase.AccountUsecase
}

func NewAccountHandler(usecase usecase.AccountUsecase) AccountHandler {
	return &accountHandler{usecase}
}

func (h accountHandler) Login(w http.ResponseWriter, r *http.Request) {

	var requestPayload struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	if err := pkg.ReadJSON(w, r, &requestPayload); err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
	}

	token, err := h.accountUsecase.Login(requestPayload.UserName, requestPayload.Password)

	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
	}

	payload := pkg.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("get account data"),
		Data:    token,
	}
	pkg.WriteJson(w, http.StatusAccepted, payload)
}
