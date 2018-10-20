package setup

import (
	"context"

	"github.com/rancher/multi-cluster-app/api/globaldns"
	"github.com/rancher/norman/store/crd"
	"github.com/rancher/norman/types"
	managementschema "github.com/rancher/types/apis/management.cattle.io/v3/schema"
	"github.com/rancher/types/client/management/v3"
	"github.com/rancher/types/config"
)

func Schemas(ctx context.Context, management *config.ManagementContext, schemas *types.Schemas) error {
	factory := &crd.Factory{ClientGetter: management.ClientGetter}

	factory.BatchCreateCRDs(ctx, config.ManagementStorageContext, schemas, &managementschema.Version,
		client.GlobalDNSType)

	factory.BatchWait()

	GlobalDNS(schemas)
	return nil
}

func GlobalDNS(schemas *types.Schemas) {
	schema := schemas.Schema(&managementschema.Version, client.GlobalDNSType)
	schema.ListHandler = globaldns.DNSListHandler
}
