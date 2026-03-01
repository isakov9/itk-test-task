package handler

import (
	"encoding/json"
	"errors"
	"itk-test-task/iternal/repository"
	"itk-test-task/iternal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handler) GetWalletById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	walletID := chi.URLParam(r, "wallet_id")

	if _, err := uuid.Parse(walletID); err != nil {
		http.Error(w, "invalid wallet id", http.StatusBadRequest)
		return
	}

	amount, err := h.walletService.GetWalletById(ctx, walletID)
	if err != nil {
		if errors.Is(err, repository.ErrWalletNotFound) {
			http.Error(w, "wallet not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(amount)
}

func (h *Handler) UpdateWallet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request UpdateWalletRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if request.Amount <= 0 {
		http.Error(w, "invalid amount", http.StatusBadRequest)
		return
	}

	if request.OperationType != service.OperationDeposit && request.OperationType != service.OperationWithdraw {
		http.Error(w, "invalid operation type", http.StatusBadRequest)
		return
	}

	input := &service.UpdateWallet{
		WalletID:      request.WalletID,
		OperationType: request.OperationType,
		Amount:        request.Amount,
	}

	err := h.walletService.UpdateWallet(ctx, input)
	if err != nil {
		if errors.Is(err, repository.ErrWalletNotFound) {
			http.Error(w, "wallet not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
