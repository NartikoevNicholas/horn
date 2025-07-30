package horn

func Run(app App, config Config) {

	server := server{app: app, config: config}
	server.startup()
}
