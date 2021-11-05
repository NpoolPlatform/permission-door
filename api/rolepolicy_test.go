package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/poolPlatform/permission-door/message/npool"
	"github.com/stretchr/testify/assert"
)

func TestSetRolePolicies(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	// set role policy correctly.
	fmt.Println("set role policy correctly.")
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.SetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "123",
			ResourcePolicies: []*npool.ResourcePolicy{
				{ResourceId: "1", Action: "get"},
				{ResourceId: "2", Action: "post"},
				{ResourceId: "3", Action: "delete"},
				{ResourceId: "4", Action: "patch"},
			},
		}).Post("http://localhost:32759/v1/set/role/policies")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.SetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "124",
			ResourcePolicies: []*npool.ResourcePolicy{
				{ResourceId: "1", Action: "get"},
				{ResourceId: "2", Action: "post"},
				{ResourceId: "3", Action: "delete"},
				{ResourceId: "4", Action: "patch"},
			},
		}).Post("http://localhost:32759/v1/set/role/policies")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.SetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "125",
			ResourcePolicies: []*npool.ResourcePolicy{
				{ResourceId: "1", Action: "get"},
				{ResourceId: "2", Action: "post"},
				{ResourceId: "3", Action: "delete"},
				{ResourceId: "4", Action: "patch"},
			},
		}).Post("http://localhost:32759/v1/set/role/policies")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	// set role policies which are already exist.
	fmt.Println("set role policies which are already exist.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.SetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "123",
			ResourcePolicies: []*npool.ResourcePolicy{
				{ResourceId: "1", Action: "get"},
				{ResourceId: "2", Action: "post"},
			},
		}).Post("http://localhost:32759/v1/set/role/policies")
	assert.Nil(t, err)
	fmt.Printf("resp status code is: %v", resp.StatusCode())

	// set role policies which are exist and not exist.
	fmt.Println("set role policies which are exist and not exist.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.SetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "123",
			ResourcePolicies: []*npool.ResourcePolicy{
				{ResourceId: "1", Action: "get"},
				{ResourceId: "5", Action: "post"},
			},
		}).Post("http://localhost:32759/v1/set/role/policies")
	assert.Nil(t, err)
	fmt.Printf("resp status code is: %v", resp.StatusCode())
}

func TestGetRolePolicies(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	response := npool.GetRolePoliciesResponse{}

	cli := resty.New()

	// get role policies with correct params.
	fmt.Println("get role policies with correct params.")
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "123",
		}).Post("http://localhost:32759/v1/get/role/policies")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &response)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response.Infos, npool.RolePolicies{})
		}
	}

	// get role policies with wrong RoleId.
	fmt.Println("get role policies with wrong RoleId.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "123",
		}).Post("http://localhost:32759/v1/get/role/policies")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		response = npool.GetRolePoliciesResponse{}
		err := json.Unmarshal(resp.Body(), &response)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response.Infos, npool.RolePolicies{})
		}
	}

	// get role policies with wrong AppId.
	fmt.Println("get role policies with wrong AppId.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "123",
		}).Post("http://localhost:32759/v1/get/role/policies")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		response = npool.GetRolePoliciesResponse{}
		err := json.Unmarshal(resp.Body(), &response)
		if assert.Nil(t, err) {
			assert.NotEqual(t, response.Infos, npool.RolePolicies{})
		}
	}
}

func TestUnsetRolePolicies(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	// normal unset policy.
	fmt.Println("normal unset policy.")
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UnsetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "123",
			Policies: []*npool.ResourcePolicy{
				{ResourceId: "1", Action: "get"},
			},
		}).Post("http://localhost:32759/v1/unset/role/policies")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	// unset policy with wrong action.
	fmt.Println("unset policy action wrong.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UnsetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "123",
			Policies: []*npool.ResourcePolicy{
				{ResourceId: "2", Action: "get"},
			},
		}).Post("http://localhost:32759/v1/unset/role/policies")
	assert.Nil(t, err)
	assert.NotEqual(t, 200, resp.StatusCode())

	// unset policy with correct role but wrong resource.
	fmt.Println("unset policy with correct role but wrong resource.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UnsetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "123",
			Policies: []*npool.ResourcePolicy{
				{ResourceId: "5", Action: "get"},
			},
		}).Post("http://localhost:32759/v1/unset/role/policies")
	assert.Nil(t, err)
	assert.NotEqual(t, 200, resp.StatusCode())

	// unset policy which has been deleted.
	fmt.Println("unset policy which has been deleted.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UnsetRolePoliciesRequest{
			AppId:  "123456789",
			RoleId: "123",
			Policies: []*npool.ResourcePolicy{
				{ResourceId: "1", Action: "get"},
			},
		}).Post("http://localhost:32759/v1/unset/role/policies")
	assert.Nil(t, err)
	assert.NotEqual(t, 200, resp.StatusCode())

	// unset policy which not exit.
	fmt.Println("unset policy which not exit.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UnsetRolePoliciesRequest{
			AppId:  "987654321",
			RoleId: "321",
			Policies: []*npool.ResourcePolicy{
				{ResourceId: "11", Action: "random"},
			},
		}).Post("http://localhost:32759/v1/unset/role/policies")
	assert.Nil(t, err)
	assert.NotEqual(t, 200, resp.StatusCode())
}

func TestDeleteRole(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	cli := resty.New()

	// delete role with wrong domain or app_id(not exist).
	fmt.Println("delete role with wrong domain or app_id(not exist).")
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteRoleRequest{
			AppId:   "1234567890",
			RoleIds: []string{"123"},
		}).Post("http://localhost:32759/v1/delete/role")
	assert.Nil(t, err)
	assert.NotEqual(t, 200, resp.StatusCode())

	// delete role with a wrong param and a correct one.
	fmt.Println("delete role with a wrong param and a correct one.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteRoleRequest{
			AppId:   "123456789",
			RoleIds: []string{"124", "12345"},
		}).Post("http://localhost:32759/v1/delete/role")
	assert.Nil(t, err)
	assert.NotEqual(t, 200, resp.StatusCode())

	// delete role with correct params.
	fmt.Println("delete role with correct params.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteRoleRequest{
			AppId:   "123456789",
			RoleIds: []string{"125"},
		}).Post("http://localhost:32759/v1/delete/role")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
