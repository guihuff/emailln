package endpoints

import (
	"bytes"
	"context"
	"emailln/internal/contract"
	internalmock "emailln/internal/test/internalMock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	body = contract.NewCampaign{
		Name:      "teste",
		Content:   "Hi everyone",
		Emails:    []string{"teste@teste.com"},
		CreatedBy: "teste@create.com",
	}
)

func setup() (*http.Request, *httptest.ResponseRecorder) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	ctx := context.WithValue(req.Context(), "email", body.CreatedBy)
	req = req.WithContext(ctx)
	res := httptest.NewRecorder()

	return req, res
}

func Test_CampaignPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name &&
			request.Content == body.Content &&
			request.CreatedBy == body.CreatedBy {
			return true
		} else {
			return false
		}
	})).Return("34x", nil)
	handler := Handler{CampaignService: service}

	req, res := setup()

	_, status, err := handler.CampaignPost(res, req)

	assert.Equal(201, status)
	assert.Nil(err)
}

func Test_CampaignPost_should_inform_error_when_exist(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}

	req, res := setup()

	_, _, err := handler.CampaignPost(res, req)

	assert.NotNil(err)
}
