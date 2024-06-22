package server

func RunServer() error {
	server, err := InitRouter()
	if err != nil {
		return err
	}
	return server.Run(":9004")
}
