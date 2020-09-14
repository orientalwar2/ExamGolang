package main

type app struct {
	config *AppConfig
}

func (a app) run() {
	NewBarangDelivery(a.config).create()
}

func newApp() app {
	config := newConfig()
	return app{config: config}
}

func main() {
	newApp().run()
}
