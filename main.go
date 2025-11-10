package main

// Exported interface
type NewPluginService interface {
	Greet(name string) string
}

// Exported struct implementing the interface
type NewPluginServiceImp struct{}

// Implementation of Greet
func (p *NewPluginServiceImp) Greet(name string) string {
	return "Hello, " + name + " from plugin!"
}

// Exported factory function returning the interface
func New() NewPluginService {
	return &NewPluginServiceImp{}
}
