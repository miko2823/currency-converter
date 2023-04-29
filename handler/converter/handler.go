package converter

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/miko2823/currency-converter.git/pkg"
	usecase "github.com/miko2823/currency-converter.git/usecase/converter"
)

type ConverterHandler interface {
	RegisterHandlers() http.Handler
	getLatestRates(w http.ResponseWriter, r *http.Request)
	getAllSymbols(w http.ResponseWriter, r *http.Request)
}

type converterHandler struct {
	converterUsecase usecase.ConverterUsecase
}

func NewConverterHandler(usecase usecase.ConverterUsecase) ConverterHandler {
	return &converterHandler{usecase}
}

func (h converterHandler) getLatestRates(w http.ResponseWriter, r *http.Request) {

	base := r.URL.Query().Get("base")
	symbols := r.URL.Query().Get("symbols")
	amount := r.URL.Query().Get("amount")
	intAmount, _ := strconv.Atoi(amount)

	latestRates, err := h.converterUsecase.GetLatestRates(base, symbols, intAmount)

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

func (h converterHandler) getAllSymbols(w http.ResponseWriter, r *http.Request) {

	latestRates, err := h.converterUsecase.GetAllSymbols()

	if err != nil {
		pkg.ErrorJSON(w, err, http.StatusBadRequest)
	}

	payload := pkg.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("get all simbols"),
		Data:    latestRates,
	}
	pkg.WriteJson(w, http.StatusAccepted, payload)
}
