package endpoints

import (
	"emailln/internal/contract"
	internalmock "emailln/internal/test/mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGetById_should_return_campaign(t *testing.T) {
	assert := assert.New(t)
	campaignResponse := contract.CampaignResponse{
		ID:      "341",
		Name:    "teste",
		Content: "Hi teste ponto com!",
		Status:  "Pending",
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaignResponse, nil)
	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(res, req)

	assert.Equal(200, status)
	assert.Equal(campaignResponse.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(campaignResponse.Name, response.(*contract.CampaignResponse).Name)

}

func Test_CampaignGetById_should_return_error_when_something_wrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)

	errExpected := errors.New("something wrong")
	service.On("GetBy", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	_, _, err := handler.CampaignGetById(res, req)

	assert.Equal(errExpected.Error(), err.Error())

}
