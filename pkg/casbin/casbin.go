package casbin

import (
	"github.com/NpoolPlatform/permission-door/pkg/db"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entadapter "github.com/casbin/ent-adapter"
	"golang.org/x/xerrors"
)

var enforcer *casbin.Enforcer

func NewCasbinEnforcer() error {
	client, err := db.Client()
	if err != nil {
		return err
	}
	adapter, err := entadapter.NewAdapterWithClient(client)
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
