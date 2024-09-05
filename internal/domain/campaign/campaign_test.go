package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email1@teste.com", "email2@teste.com"}
)

func Test_NewCampaign(t *testing.T) {
	//arrange
	assert := assert.New(t)
	timeNow := time.Now()

	//action
	campaign, _ := NewCampaign(name, content, contacts)

	//assert
	assert.NotNil(campaign.ID)
	assert.GreaterOrEqual(campaign.CreatedOn, timeNow)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_ShouldReturnAnErrorWhenNameIsEmpty(t *testing.T) {
	//arrange
	assert := assert.New(t)
	emptyName := ""

	//action
	_, err := NewCampaign(emptyName, content, contacts)

	//assert
	assert.NotNil(err)
	assert.Equal("the name is required", err.Error())
}

func Test_NewCampaign_ShouldReturnAnErrorWhenContentIsEmpty(t *testing.T) {
	//arrange
	assert := assert.New(t)
	emptyContent := ""

	//action
	_, err := NewCampaign(name, emptyContent, contacts)

	//assert
	assert.NotNil(err)
	assert.Equal("the content is required", err.Error())
}

func Test_NewCampaign_ShouldReturnAnErrorWhenContactsIsEmpty(t *testing.T) {
	//arrange
	assert := assert.New(t)

	//action
	_, err := NewCampaign(name, content, nil)

	//assert
	assert.NotNil(err)
	assert.Equal("the contacts is required", err.Error())
}
