package api

import (
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/permission-door/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	userID := uuid.New().String()
	appID := uuid.New().String()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.AuthenticateUserPolicyByIDRequest{
			UserID: userID,
			AppID:  appID,
		}).Post("http://localhost:50100/v1/authenticate/user/policy/by/id")
	if assert.Nil(t, err) {
		assert.NotEqual(t, 200, resp.StatusCode())
	}
}
