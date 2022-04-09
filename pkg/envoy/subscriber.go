// @since 2022-04-09 19:17:37
// @author acrazing <joking.young@gmail.com>

package envoy

import "github.com/acrazing/universal-ingress-controller/pkg/core"

const name = "envoy"

type envoySubscriber struct {
}

func NewEnvoySubscriber(configFile string) core.Subscriber {
	return &envoySubscriber{}
}

func init() {
	core.RegisterSubscriber(name, NewEnvoySubscriber)
}
