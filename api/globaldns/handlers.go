package globaldns

import (
	"github.com/rancher/norman/api/handler"
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
)

func DNSListHandler(request *types.APIContext, next types.RequestHandler) error {
	logrus.Debugf("TokenListHandler called")

	return handler.ListHandler(request, next)
}
