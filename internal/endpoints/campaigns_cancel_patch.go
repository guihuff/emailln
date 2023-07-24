package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignCancelPatch(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	paramId := chi.URLParam(r, "id")
	err := h.CampaignService.Cancel(paramId)
	return nil, 200, err
}
