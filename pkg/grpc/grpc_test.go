package grpc

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"
	testinit "github.com/poolPlatform/permission-door/pkg/test-init"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestGrpc(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userID := uuid.New().String()
	appID := uuid.New().String()
	_, err := GetUserRoleIDs(userID, appID)
	assert.NotNil(t, err)
}
