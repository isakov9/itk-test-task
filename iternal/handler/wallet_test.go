package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"itk-test-task/iternal/handler"
	"itk-test-task/iternal/mocks"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter(h *handler.Handler) chi.Router {
	r := chi.NewRouter()
	r.Mount("/api/v1", h.NewRouter())
	return r
}

func TestUpdateWallet_InvalidAmount(t *testing.T) {
	mockService := mocks.NewWalletServiceInterface(t)
	h := handler.NewHandler(mockService)

	body, _ := json.Marshal(handler.UpdateWalletRequest{
		WalletID:      "123e4567-e89b-12d3-a456-426614174000",
		OperationType: "DEPOSIT",
		Amount:        -100,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	setupRouter(h).ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertNotCalled(t, "UpdateWallet")
}

func TestUpdateWallet_InvalidOperationType(t *testing.T) {
	mockService := mocks.NewWalletServiceInterface(t)
	h := handler.NewHandler(mockService)

	body, _ := json.Marshal(handler.UpdateWalletRequest{
		WalletID:      "123e4567-e89b-12d3-a456-426614174000",
		OperationType: "INVALID",
		Amount:        100,
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	setupRouter(h).ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetWalletById_InvalidUUID(t *testing.T) {
	mockService := mocks.NewWalletServiceInterface(t)
	h := handler.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets/not-a-uuid", nil)
	w := httptest.NewRecorder()

	setupRouter(h).ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertNotCalled(t, "GetWalletById")
}

func TestGetWalletById_Success(t *testing.T) {
	mockService := mocks.NewWalletServiceInterface(t)
	mockService.On("GetWalletById", mock.Anything, "123e4567-e89b-12d3-a456-426614174000").Return(int64(1000), nil)

	h := handler.NewHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets/123e4567-e89b-12d3-a456-426614174000", nil)
	w := httptest.NewRecorder()

	setupRouter(h).ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
