package api

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/permission-door/message/npool"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteResource(t *testing.T) {
	fmt.Println("start delete resource")
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	appID := uuid.New().String()
	roleID := uuid.New().String()

	fmt.Println("start set policy")
	cli := resty.New()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.SetRolePoliciesRequest{
			AppId:  appID,
			RoleId: roleID,
			ResourcePolicies: []*npool.ResourcePolicy{
				{ResourceId: "1", Action: "get"},
				{ResourceId: "2", Action: "post"},
				{ResourceId: "3", Action: "delete"},
				{ResourceId: "4", Action: "patch"},
			},
		}).Post("http://localhost:50100/v1/set/role/policies")
	if assert.Nil(t, err) {
		fmt.Println("set resp is", resp)
		assert.Equal(t, 200, resp.StatusCode())
	}

	// delete resource with correct params.
	fmt.Println("delete resource with correct params.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteResourceRequest{
			AppId:       appID,
			ResourceIds: []string{"1"},
		}).Post("http://localhost:50100/v1/delete/resource")
	if assert.Nil(t, err) {
		fmt.Println("delete resp is", resp)
		assert.Equal(t, 200, resp.StatusCode())
	}
}
