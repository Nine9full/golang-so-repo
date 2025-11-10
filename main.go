package main

import pluginiface "github.com/Nine9full/golang-so-repo/plugin-interface"

// Exported interface

// Exported struct implementing the interface
type NewPluginServiceImp struct{}

// Implementation of Greet
func (p *NewPluginServiceImp) Greet(name string) string {
	return "Hello, " + name + " from plugin!"
}

// Exported factory function returning the interface
func New() pluginiface.NewPluginService {
	return &NewPluginServiceImp{}
}
