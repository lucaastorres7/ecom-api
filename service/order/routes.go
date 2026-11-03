package cart

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucaastorres7/ecom-api/types"
)

type Handler struct {
	store types.CartStore
}

func NewHandler(store types.CartStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {

}