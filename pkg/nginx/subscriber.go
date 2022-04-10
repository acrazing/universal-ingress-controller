// @since 2022-04-09 19:17:37
// @author acrazing <joking.young@gmail.com>

package nginx

import (
	"github.com/acrazing/universal-ingress-controller/pkg/core"
	"google.golang.org/grpc"
)

const name = "nginx"

type nginxSubscriber struct {
}

func (n nginxSubscriber) TransformIngresses(ingresses map[string]*core.Ingress) error {
	//TODO implement me
	panic("implement me")
}

func (n nginxSubscriber) UpdateResources(resources map[string]*core.SubscribeResource) error {
	//TODO implement me
	panic("implement me")
}

func (n nginxSubscriber) RegisterGrpcServers(s *grpc.Server) {
	//TODO implement me
	panic("implement me")
}

func NewNginxSubscriber() core.Subscriber {
	return &nginxSubscriber{}
}

func init() {
	core.RegisterSubscriber(name, NewNginxSubscriber)
}
