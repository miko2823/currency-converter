package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/miko2823/currency-converter.git/pkg"
	"github.com/miko2823/currency-converter.git/usecase"
)

type ConverterHandler interface {
	RegisterHandlers() http.Handler
	getLatestRates(w http.ResponseWriter, r *http.Request)
}

type converterHandler struct {
	converterUsecase usecase.ConverterUsecase
}

func NewConverterHandler(usecase usecase.ConverterUsecase) ConverterHandler {
	return &converterHandler{usecase}
}

func (p converterHandler) getLatestRates(w http.ResponseWriter, r *http.Request) {

	// var requestPayload struct {
	// 	Id string `json:"id"`
	// }
	// err := pkg.ReadJSON(w, r, &requestPayload)
	// fmt.Println(err)
	// if err != nil {
	// 	pkg.ErrorJSON(w, err, http.StatusBadRequest)
	// }
	base := r.URL.Query().Get("base")
	symbols := r.URL.Query().Get("symbols")
	amount := r.URL.Query().Get("amount")
	intAmount, _ := strconv.Atoi(amount)

	latestRates, err := p.converterUsecase.GetLatestRates(base, symbols, intAmount)

	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
	}

	payload := pkg.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("get converter data"),
		Data:    latestRates,
	}
	pkg.WriteJson(w, http.StatusAccepted, payload)
}
