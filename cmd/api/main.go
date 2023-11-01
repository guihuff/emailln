package main

import (
	"log"
	"os"

	"emailln/internal/domain/campaign"
	"emailln/internal/endpoints"
	"emailln/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDB()
	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)

		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
		r.Delete("/delete/{id}", endpoints.HandlerError(handler.CampaignDelete))
	})

	log.Print("API listening port: " + os.Getenv("API_PORT"))
	http.ListenAndServe(":"+os.Getenv("API_PORT"), r)

}
