package user

import (
	"context"

	"github.com/NpoolPlatform/permission-door/message/npool"
	"github.com/NpoolPlatform/permission-door/pkg/casbin"
	"github.com/NpoolPlatform/permission-door/pkg/grpc"
	"golang.org/x/xerrors"
)

func AuthenticateUserPolicyByID(ctx context.Context, in *npool.AuthenticateUserPolicyByIDRequest) (*npool.AuthenticateUserPolicyByIDResponse, error) {
	resp, err := grpc.GetUserRoleIDs(ctx, in.UserID, in.AppID)
	if err != nil {
		return nil, xerrors.Errorf("fail to get user role: %v", err)
	}

	for _, role := range resp {
		ok, err := casbin.Enforcer().Enforce(role, in.AppID, in.ResourecID, in.Action)
		if err != nil || !ok {
			continue
		}
		return &npool.AuthenticateUserPolicyByIDResponse{
			Info: "ok",
		}, nil
	}
	return nil, xerrors.Errorf("user has no permission")
}
