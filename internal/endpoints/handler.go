package endpoints

import "emailln/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}

