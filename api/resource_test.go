package api

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/poolPlatform/permission-door/message/npool"
	"github.com/stretchr/testify/assert"
)

func TestDeleteResource(t *testing.T) {
	fmt.Println("start delete resource")
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	fmt.Println("start set policy")
	cli := resty.New()
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

	// delete resource with wrong ResourceId.
	fmt.Println("delete resource with wrong ResourceId.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteResourceRequest{
			AppId:       "123456789",
			ResourceIds: []string{"12345"},
		}).Post("http://localhost:32759/v1/delete/resource")
	assert.Nil(t, err)
	assert.NotEqual(t, 200, resp.StatusCode())

	// delete resource with wrong AppId.
	fmt.Println("delete resource with wrong AppId.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteResourceRequest{
			AppId:       "1234567890",
			ResourceIds: []string{"1"},
		}).Post("http://localhost:32759/v1/delete/resource")
	assert.Nil(t, err)
	assert.NotEqual(t, 200, resp.StatusCode())

	// delete resource with wrong params.
	fmt.Println("delete resource with wrong params.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteResourceRequest{
			AppId:       "1234567890",
			ResourceIds: []string{"12345"},
		}).Post("http://localhost:32759/v1/delete/resource")
	assert.Nil(t, err)
	assert.NotEqual(t, 200, resp.StatusCode())

	// delete resource with correct params.
	fmt.Println("delete resource with correct params.")
	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteResourceRequest{
			AppId:       "123456789",
			ResourceIds: []string{"1"},
		}).Post("http://localhost:32759/v1/delete/resource")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
