package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/poolPlatform/permission-door/message/npool"
	"github.com/stretchr/testify/assert"
)

func TestSetRolePolicies(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()
	appID := uuid.New().String()
	roleID := uuid.New().String()
	// set role policy correctly.
	fmt.Println("set role policy correctly.")
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
	assert.Nil(t, err)
	fmt.Print(resp)
	assert.Equal(t, 200, resp.StatusCode())

	response := npool.GetRolePoliciesResponse{}

	// get role policies with correct params.
	fmt.Println("get role policies with correct params.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetRolePoliciesRequest{
			AppId:  appID,
			RoleId: roleID,
		}).Post("http://localhost:50100/v1/get/role/policies")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &response)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response.Infos, npool.RolePolicies{})
		}
	}

	// normal unset policy.
	fmt.Println("normal unset policy.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UnsetRolePoliciesRequest{
			AppId:  appID,
			RoleId: roleID,
			Policies: []*npool.ResourcePolicy{
				{ResourceId: "1", Action: "get"},
			},
		}).Post("http://localhost:50100/v1/unset/role/policies")
	assert.Nil(t, err)
	fmt.Println("unset role policies resp is", resp)
	assert.Equal(t, 200, resp.StatusCode())

	// delete role with correct params.
	fmt.Println("delete role with correct params.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteRoleRequest{
			AppId:   appID,
			RoleIds: []string{roleID},
		}).Post("http://localhost:50100/v1/delete/role")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
