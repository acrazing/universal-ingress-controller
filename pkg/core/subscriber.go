// @since 2022-04-09 19:41:54
// @author acrazing <joking.young@gmail.com>

package core

import (
	"fmt"
	"google.golang.org/grpc"
)

type SubscribeResource struct {
	Name     string
	Revision int
	Filename string
	Data     []byte
}

type Subscriber interface {
	TransformIngresses(ingresses map[string]*Ingress) error
	UpdateResources(resources map[string]*SubscribeResource) error
	RegisterGrpcServers(s *grpc.Server)
}

type SubscriberFactory func() Subscriber

var subscriberFactories = map[string]SubscriberFactory{}

func RegisterSubscriber(name string, factory SubscriberFactory) {
	if _, ok := subscriberFactories[name]; ok {
		panic(fmt.Sprintf("subscriber `%s` is registered already.", name))
	}
	subscriberFactories[name] = factory
}
