package handlers

import (
	cartdto "backend/dto/cart"
	dto "backend/dto/result"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var cart models.Cart
	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) CreateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(cartdto.CreateCartRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	status := "pending"

	cart := models.Cart{
		Status: status,
		Qty:    request.Qty,
		UserID: request.UserID,
	}

	cart, err = h.CartRepository.CreateCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	cart, _ = h.CartRepository.GetCart(cart.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cart}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) UpdateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(cartdto.UpdateCartRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	cart, err := h.CartRepository.GetCart(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "aa.Error()"}
		json.NewEncoder(w).Encode(response)
		return
	}
	if request.Qty != 0 {
		cart.Qty = request.Qty
	}
	if request.SubTotal != 0 {
		cart.SubTotal = request.SubTotal
	}
	if request.UserID != 0 {
		cart.UserID = request.UserID
	}
	if request.Status != "" {
		cart.Status = request.Status
	}

	data, err := h.CartRepository.UpdateCart(cart, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
