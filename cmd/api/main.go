package main

import (
	"emailln/internal/domain/campaign"
	"emailln/internal/endpoints"
	"emailln/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Hendler{
		CampaignService: campaignService,
	}

	r.Post("/campaigns", handler.CampaignPost)
	r.Get("/campaigns", handler.CampaignGet)


	println("Porta: 3000")
	http.ListenAndServe(":3000", r)

}
