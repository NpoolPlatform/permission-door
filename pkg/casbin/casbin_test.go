package casbin

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/permission-door/pkg/test-init"
)

func init() {
	fmt.Println("start init")
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	if err := testinit.Init(); err != nil {
		log.Fatalf("cannot init test stub: %v\n", err)
	}
}

func TestNewCasbinEnforcer(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	err := NewCasbinEnforcer()
	if err != nil {
		fmt.Printf("new casbin enforcer error: %v", err)
		return
	}
}
