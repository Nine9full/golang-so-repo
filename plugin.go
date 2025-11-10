package main

import (
	iface "github.com/Nine9full/golang-main-repo/plugin-interface"
)

// Exported interface

// Exported struct implementing the interface
type NewPluginServiceImp struct{}

// Exported factory function returning the interface
func New() iface.NewPluginService {
	return &NewPluginServiceImp{}
}

// Implementation of Greet
func (p *NewPluginServiceImp) Greet(name string) string {
	return "Hello, " + name + " from plugin!"
}
