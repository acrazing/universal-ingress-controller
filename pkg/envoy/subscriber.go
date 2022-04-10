// @since 2022-04-09 19:17:37
// @author acrazing <joking.young@gmail.com>

package envoy

import (
	envoypb "github.com/acrazing/universal-ingress-controller/gen/envoy"
	"github.com/acrazing/universal-ingress-controller/pkg/core"
	discoveryv3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"google.golang.org/grpc"
)

const name = "envoy"

type envoySubscriber struct {
	config *envoypb.EnvoyConfig
	ads    *ads
}

func (e *envoySubscriber) TransformIngresses(ingresses map[string]*core.Ingress) error {
	//TODO implement me
	panic("implement me")
}

func (e *envoySubscriber) UpdateResources(resources map[string]*core.SubscribeResource) error {
	//TODO implement me
	panic("implement me")
}

func (e *envoySubscriber) RegisterGrpcServers(s *grpc.Server) {
	discoveryv3.RegisterAggregatedDiscoveryServiceServer(s, e.ads)
}

func NewEnvoySubscriber() core.Subscriber {
	return &envoySubscriber{
		config: &envoypb.EnvoyConfig{},
		ads:    newAds(),
	}
}

func init() {
	core.RegisterSubscriber(name, NewEnvoySubscriber)
}
