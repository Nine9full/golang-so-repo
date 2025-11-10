package main

// HiPlugin implements Greeter
type NewPluginServiceImp struct{}

type NewPluginService interface {
	Greet(name string) string
}

// New returns the Greeter interface
func New() NewPluginService {
	return &NewPluginServiceImp{}
}

func (h *NewPluginServiceImp) Greet(name string) string {
	return "Hi, " + name + " from NewPluginService!"
}
