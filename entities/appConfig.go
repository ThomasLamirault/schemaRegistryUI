package entities

type AppConfig struct {
	Registry Registry
	App      App
}

type Registry struct {
	Address string
	Port    string
}

type App struct {
	Address string
	Port    int
}
