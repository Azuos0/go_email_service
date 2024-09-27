package campaign

import (
	"go_email_service/internal/contracts"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	//arrange
	assert := assert.New(t)
	newCampaign := contracts.NewCampaign{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"teste1@teste.com"},
	}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		}
		if campaign.Content != newCampaign.Content {
			return false
		}
		if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service := Service{Repository: repositoryMock}

	//act
	id, err := service.Create(newCampaign)

	//assert
	assert.NotNil(id)
	repositoryMock.AssertExpectations(t)
	assert.Nil(err)
}
