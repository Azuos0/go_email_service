package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	if err := validateCampaignRequiredFields(name, content, emails); err != nil {
		return nil, err
	}

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index] = Contact{
			Email: email,
		}
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}, nil
}

func validateCampaignRequiredFields(name string, content string, emails []string) error {
	if name == "" {
		return errors.New("the name is required")
	}

	if content == "" {
		return errors.New("the content is required")
	}

	if len(emails) == 0 {
		return errors.New("the contacts is required")
	}

	return nil
}
