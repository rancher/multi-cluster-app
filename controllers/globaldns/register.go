package globaldns

import (
	"context"

	"github.com/rancher/types/config"
)

func Register(ctx context.Context, management *config.ManagementContext) {
	n := newGlobalDNSController(management)
	management.Management.GlobalDNSs("").AddHandler(GlobaldnsController, n.sync)
}
