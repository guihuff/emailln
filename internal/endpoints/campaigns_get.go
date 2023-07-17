package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Hendler) CampaignGet(w http.ResponseWriter, r *http.Request) {

	campaigns, _ := h.CampaignService.Repository.Get()
	render.Status(r, 200)
	render.JSON(w, r, campaigns)
}
