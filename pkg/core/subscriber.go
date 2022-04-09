// @since 2022-04-09 19:41:54
// @author acrazing <joking.young@gmail.com>

package core

import "fmt"

type Subscriber interface {
}

type SubscriberFactory func(configFile string) Subscriber

var subscriberFactories = map[string]SubscriberFactory{}

func RegisterSubscriber(name string, factory SubscriberFactory) {
	if _, ok := subscriberFactories[name]; ok {
		panic(fmt.Sprintf("subscriber `%s` is registered already.", name))
	}
	subscriberFactories[name] = factory
}
