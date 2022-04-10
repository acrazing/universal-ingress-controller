// @since 2022-04-09 19:17:29
// @author acrazing <joking.young@gmail.com>

package envoy

import discoveryv3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"

type ads struct {
}

func (a ads) StreamAggregatedResources(server discoveryv3.AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	//TODO implement me
	panic("implement me")
}

func (a ads) DeltaAggregatedResources(server discoveryv3.AggregatedDiscoveryService_DeltaAggregatedResourcesServer) error {
	//TODO implement me
	panic("implement me")
}

func newAds() *ads {
	return &ads{}
}
