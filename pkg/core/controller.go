// @since 2022-04-09 19:44:14
// @author acrazing <joking.young@gmail.com>

package core

type UniversalIngressController struct {
	subscribers []Subscriber
}

func (c *UniversalIngressController) Run() error {
	panic("not implemented")
}

func NewUniversalIngressController(configFile string) *UniversalIngressController {
	return &UniversalIngressController{}
}
