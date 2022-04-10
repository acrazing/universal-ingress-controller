// @since 2022-04-09 19:17:37
// @author acrazing <joking.young@gmail.com>

package envoy

import (
	"context"
	envoyv1 "github.com/acrazing/universal-ingress-controller/gen/envoy_subscriber/v1"
	"github.com/acrazing/universal-ingress-controller/pkg/core"
	"github.com/envoyproxy/go-control-plane/envoy/service/cluster/v3"
	"github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/envoy/service/endpoint/v3"
	"github.com/envoyproxy/go-control-plane/envoy/service/listener/v3"
	"github.com/envoyproxy/go-control-plane/envoy/service/route/v3"
	"github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	"github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/log"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"google.golang.org/grpc"
)

const (
	name = "envoy"
)

type envoySubscriber struct {
	config *envoyv1.EnvoyConfig
	cache  cache.SnapshotCache
	srv    server.Server
}

func (e *envoySubscriber) TransformIngresses(ingresses map[string]*core.Ingress) error {
	//TODO implement me
	return nil
}

func (e *envoySubscriber) UpdateResources(resources map[string]*core.SubscribeResource) error {

	panic("implement me")
}

func (e *envoySubscriber) RegisterGrpcServers(s *grpc.Server) {
	envoy_service_discovery_v3.RegisterAggregatedDiscoveryServiceServer(s, e.srv)
	envoy_service_endpoint_v3.RegisterEndpointDiscoveryServiceServer(s, e.srv)
	envoy_service_cluster_v3.RegisterClusterDiscoveryServiceServer(s, e.srv)
	envoy_service_route_v3.RegisterRouteDiscoveryServiceServer(s, e.srv)
	envoy_service_listener_v3.RegisterListenerDiscoveryServiceServer(s, e.srv)
	envoy_service_secret_v3.RegisterSecretDiscoveryServiceServer(s, e.srv)
	envoy_service_runtime_v3.RegisterRuntimeDiscoveryServiceServer(s, e.srv)
}

func NewEnvoySubscriber() core.Subscriber {
	c := cache.NewSnapshotCache(false, cache.IDHash{}, log.NewDefaultLogger())
	srv := server.NewServer(context.Background(), c, server.CallbackFuncs{})
	return &envoySubscriber{
		config: &envoyv1.EnvoyConfig{},
		cache:  c,
		srv:    srv,
	}
}

func init() {
	core.RegisterSubscriber(name, NewEnvoySubscriber)
}
