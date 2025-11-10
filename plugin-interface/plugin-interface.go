package pluginiface

type NewPluginService interface {
	Greet(name string) string
}
