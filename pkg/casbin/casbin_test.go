package casbin

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/poolPlatform/permission-door/pkg/test-init"
)

func init() {
	fmt.Println("start init")
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestGetMysqlConfig(t *testing.T) {
	fmt.Println("start testing")
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	mysqlConfig, err := GetMysqlConfig()
	if err != nil {
		fmt.Printf("get mysql config error: %v", err)
		return
	}
	fmt.Printf("mysql config is: %v", mysqlConfig)
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
