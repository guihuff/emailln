package campaign_test

import (
	"emailln/internal/contract"
	"emailln/internal/domain/campaign"
	internalerrors "emailln/internal/internalErrors"
	internalmock "emailln/internal/test/internalMock"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi",
		Emails:  []string{"email@e.com"},
		CreatedBy: "teste@created.com.br",
	}
	service = campaign.ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Create", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetById_ReturnCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaign(
		newCampaign.Name,
		newCampaign.Content,
		newCampaign.Emails,
		newCampaign.CreatedBy,
	)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)
	service.Repository = repositoryMock

	campaignReturned, _ := service.GetBy(campaign.ID)

	assert.Equal(campaign.ID, campaignReturned.ID)
	assert.Equal(campaign.Name, campaignReturned.Name)
	assert.Equal(campaign.Content, campaignReturned.Content)
	assert.Equal(campaign.Status, campaignReturned.Status)
	assert.Equal(campaign.CreatedBy, campaignReturned.CreatedBy)

}

func Test_GetById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	assert := assert.New(t)
  
	campaign, _ := campaign.NewCampaign(
		newCampaign.Name,
		newCampaign.Content,
		newCampaign.Emails,
		newCampaign.CreatedBy,
	)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("something wrong"))
	service.Repository = repositoryMock

	_, err := service.GetBy(campaign.ID)

	assert.True(errors.Is(internalerrors.ErrInternal, err))

}

func Test_GetById_ReturnRecordNotFound_when_campaign_does_not_exist(t *testing.T) {
	assert := assert.New(t)

	campaignIDInvalid := "invalid"
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
	service.Repository = repositoryMock

	_, err := service.GetBy(campaignIDInvalid)

	assert.True(errors.Is(gorm.ErrRecordNotFound, err))

}

func Test_Delete_ReturnRecordNotFound_when_campaign_does_not_exist(t *testing.T) {
	assert := assert.New(t)

	campaignIDInvalid := "invalid"
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
	service.Repository = repositoryMock

	err := service.Delete(campaignIDInvalid)

	assert.True(errors.Is(gorm.ErrRecordNotFound, err))

}

func Test_Delete_ReturnStatusInvalid_when_campaign_status_not_equals_pending(t *testing.T) {
	assert := assert.New(t)

	campaign := &campaign.Campaign{
		ID: "1", Status: campaign.Started,
	}
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(campaign, nil)
	service.Repository = repositoryMock

	err := service.Delete(campaign.ID)

	assert.Equal("Campaign status invalid", err.Error())

}

func Test_Delete_ReturnInternalError_when_delete_has_problem(t *testing.T) {
	assert := assert.New(t)

	campaignFound, _ := campaign.NewCampaign(
		"Teste Delete", "Body test delete", []string{"teste@t.com"}, newCampaign.CreatedBy,
	)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignFound == campaign
	})).Return(errors.New("error to delete campaign"))

	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())

}

func Test_Delete_ReturnNil_when_delete_has_sucess(t *testing.T) {
	assert := assert.New(t)
  
	//TODO: fixed arg email
	campaignFound, _ := campaign.NewCampaign(
		"Teste Delete", "Body test delete", []string{"teste@t.com"}, newCampaign.CreatedBy,
	)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignFound == campaign
	})).Return(nil)

	service.Repository = repositoryMock

	err := service.Delete(campaignFound.ID)

	assert.Nil(err)

}