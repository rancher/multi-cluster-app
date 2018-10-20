package globaldns

import (
	"strings"

	"github.com/rancher/norman/types"
	"github.com/rancher/types/apis/management.cattle.io/v3"
	"github.com/rancher/types/config"
	"github.com/sirupsen/logrus"
)

const (
	GlobaldnsController = "mgmt-global-dns-controller"
)

type GlobalDNSController struct {
	Schemas            *types.Schemas
	globalDNSInterface v3.GlobalDNSInterface
	globalDNSLister    v3.GlobalDNSLister
}

func newGlobalDNSController(mgmt *config.ManagementContext) *GlobalDNSController {
	n := &GlobalDNSController{
		globalDNSInterface: mgmt.Management.GlobalDNSs(""),
		globalDNSLister:    mgmt.Management.GlobalDNSs("").Controller().Lister(),
	}
	return n
}

//sync is called periodically and on real updates
func (n *GlobalDNSController) sync(key string, obj *v3.GlobalDNS) error {
	logrus.Debugf("GlobalDNSController called")
	if obj == nil || obj.DeletionTimestamp != nil {
		return nil
	}

	if !strings.HasSuffix(obj.DNSName, ".rancher-test.com.") {
		newObj := obj.DeepCopy()
		newObj.DNSName = newObj.DNSName + ".rancher-test.com."
		if _, err := n.globalDNSInterface.Update(newObj); err != nil {
			return err
		}
	}
	return nil
}
