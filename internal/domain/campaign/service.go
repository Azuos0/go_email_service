package campaign

import (
	"go_email_service/internal/contracts"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contracts.NewCampaign) (string, error) {
	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	s.Repository.Save(campaign)

	return campaign.ID, nil
}
