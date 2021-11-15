package resource

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/poolPlatform/permission-door/message/npool"
	"github.com/poolPlatform/permission-door/pkg/casbin"
)

func DeleteResource(ctx context.Context, in *npool.DeleteResourceRequest) (*npool.DeleteResourceResponse, error) {
	for _, resource := range in.ResourceIds {
		ok, err := casbin.Enforcer().RemoveFilteredPolicy(1, in.AppId, resource)
		if err != nil || !ok {
			logger.Sugar().Errorf("delete resource error: %v", err)
			return nil, err
		}
	}
	return &npool.DeleteResourceResponse{
		Info: "successfully",
	}, nil
}
