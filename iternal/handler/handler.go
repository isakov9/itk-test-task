package handler

import (
	"itk-test-task/iternal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	walletService service.WalletServiceInterface
}

func NewHandler(walletService service.WalletServiceInterface) *Handler {
	return &Handler{walletService: walletService}
}

func (h *Handler) NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/wallets/{wallet_id}", h.GetWalletById)
	r.Post("/wallets", h.UpdateWallet)
	return r
}
