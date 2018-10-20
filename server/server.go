package server

import (
	"context"
	"net/http"

	"github.com/rancher/multi-cluster-app/api/setup"
	"github.com/rancher/multi-cluster-app/controllers"
	normanapi "github.com/rancher/norman/api"
	"github.com/rancher/norman/types"
	managementSchema "github.com/rancher/types/apis/management.cattle.io/v3/schema"
	"github.com/rancher/types/config"
)

func NewMultiClusterAppServer(ctx context.Context, management *config.ManagementContext) (http.Handler, error) {

	schemas := types.NewSchemas().AddSchemas(managementSchema.MultiClusterAppSchemas)
	if err := setup.Schemas(ctx, management, schemas); err != nil {
		return nil, err
	}

	server := normanapi.NewAPIServer()
	if err := server.AddSchemas(schemas); err != nil {
		return nil, err
	}

	controllers.Register(ctx, management)
	if err := management.Start(ctx); err != nil {
		panic(err)
	}

	return server, nil
}
