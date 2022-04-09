// @since 2022-04-09 19:17:37
// @author acrazing <joking.young@gmail.com>

package nginx

import "github.com/acrazing/universal-ingress-controller/pkg/core"

const name = "nginx"

type nginxSubscriber struct {
}

func NewNginxSubscriber(configFile string) core.Subscriber {
	return &nginxSubscriber{}
}

func init() {
	core.RegisterSubscriber(name, NewNginxSubscriber)
}
