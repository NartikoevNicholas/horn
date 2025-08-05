package horn

func Run(config Config) {
	server := Server{config: config}
	server.Init()
	server.Startup()
}
