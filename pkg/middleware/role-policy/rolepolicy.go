package rolepolicy

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/permission-door/message/npool"
	"github.com/NpoolPlatform/permission-door/pkg/casbin"
	"golang.org/x/xerrors"
)

func SetRolePolicies(ctx context.Context, in *npool.SetRolePoliciesRequest) (*npool.SetRolePoliciesResponse, error) {
	rules := [][]string{}
	for _, policy := range in.ResourcePolicies {
		rule := []string{in.RoleId, in.AppId, policy.ResourceId, policy.Action}
		rules = append(rules, rule)
	}

	ok, err := casbin.Enforcer().AddPolicies(rules)
	if err != nil || !ok {
		return nil, xerrors.Errorf("fail to set role policies: %v", err)
	}
	return &npool.SetRolePoliciesResponse{
		Info: "set policies success!",
	}, nil
}

func GetRolePolicies(ctx context.Context, in *npool.GetRolePoliciesRequest) (*npool.GetRolePoliciesResponse, error) {
	policies := casbin.Enforcer().GetFilteredPolicy(0, in.RoleId, in.AppId)

	npoolPolicy := []*npool.ResourcePolicy{}
	for _, policy := range policies {
		npoolPolicy = append(npoolPolicy, &npool.ResourcePolicy{
			ResourceId: policy[2],
			Action:     policy[3],
		})
	}

	return &npool.GetRolePoliciesResponse{
		Infos: &npool.RolePolicies{
			RoleId:   in.RoleId,
			AppId:    in.AppId,
			Policies: npoolPolicy,
		},
	}, nil
}

func AuthenticateRolePolicy(ctx context.Context, in *npool.AuthenticateRolePolicyRequest) (*npool.AuthenticateRolePolicyResponse, error) {
	ok, err := casbin.Enforcer().Enforce(in.RoleId, in.AppId, in.Policy.ResourceId, in.Policy.Action)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, err
	}
	return &npool.AuthenticateRolePolicyResponse{
		Info: "ok",
	}, nil
}

func AuthenticateRolesPolicy(ctx context.Context, in *npool.AuthenticateRolesPolicyRequest) (*npool.AuthenticateRolesPolicyResponse, error) {
	for _, role := range in.RoleIds {
		ok, err := casbin.Enforcer().Enforce(role, in.AppId, in.Policy.ResourceId, in.Policy.Action)
		if err != nil || !ok {
			continue
		}

		return &npool.AuthenticateRolesPolicyResponse{
			Info: fmt.Sprintf("role %v has the permission", role),
		}, nil
	}
	return nil, xerrors.Errorf("roles doesn't have the permission")
}

func UnsetRolePolicies(ctx context.Context, in *npool.UnsetRolePoliciesRequest) (*npool.UnsetRolePoliciesResponse, error) {
	rules := [][]string{}
	for _, policy := range in.Policies {
		rule := []string{in.RoleId, in.AppId, policy.ResourceId, policy.Action}
		rules = append(rules, rule)
	}
	ok, err := casbin.Enforcer().RemovePolicies(rules)
	if err != nil || !ok {
		return nil, xerrors.Errorf("fail to unset role policies: %v", err)
	}
	return &npool.UnsetRolePoliciesResponse{
		Info: "unset successfully",
	}, nil
}

func DeleteRole(ctx context.Context, in *npool.DeleteRoleRequest) (*npool.DeleteRoleResponse, error) {
	for _, role := range in.RoleIds {
		ok, err := casbin.Enforcer().RemoveFilteredPolicy(0, role, in.AppId)
		if err != nil || !ok {
			return nil, xerrors.Errorf("fail to delete role: %v", err)
		}
	}
	return &npool.DeleteRoleResponse{
		Info: "delete role successfully",
	}, nil
}
