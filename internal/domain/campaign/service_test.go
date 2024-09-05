package campaign

import (
	"go_email_service/internal/contracts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Campaign(t *testing.T) {
	//arrange
	assert := assert.New(t)
	service := Service{}
	newCampaign := contracts.NewCampaign{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"teste1@teste.com"},
	}

	//act
	id, err := service.Create(newCampaign)

	//assert
	assert.NotNil(id)
	assert.Nil(err)
}
