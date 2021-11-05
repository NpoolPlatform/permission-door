package casbin

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	constant "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"golang.org/x/xerrors"
)

var enforcer *casbin.Enforcer

const (
	CasbinTableName = "casbin_rules"
)

func GetMysqlConfig() (string, error) {
	service, err := config.PeekService(constant.MysqlServiceName)
	if err != nil {
		return "", xerrors.Errorf("Fail to query mysql service: %v", err)
	}

	username := config.GetStringValueWithNameSpace(constant.MysqlServiceName, "username")
	password := config.GetStringValueWithNameSpace(constant.MysqlServiceName, "password")
	myServiceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	dbname := config.GetStringValueWithNameSpace(myServiceName, "database_name")

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, service.Address, service.Port, dbname), nil
}

func NewCasbinEnforcer() error {
	dataSource, err := GetMysqlConfig()
	if err != nil {
		return xerrors.Errorf("fail to get mysql config, error: %v", err)
	}
	adapter, err := xormadapter.NewAdapterWithTableName("mysql", dataSource, CasbinTableName, true)
	if err != nil {
		return xerrors.Errorf("fail to create adapter, error: %v", err)
	}

	m, err := model.NewModelFromString(`
	[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.dom == p.dom && r.act == p.act
	`)
	if err != nil {
		return xerrors.Errorf("Fail to new model from string, error: %v", err)
	}
	enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		return xerrors.Errorf("fail to new enforcer, error: %v", err)
	}
	return err
}

func Enforcer() *casbin.Enforcer {
	return enforcer
}
