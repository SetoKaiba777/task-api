package main

import "go-challenger/infrastructure/http/server"

func main() {
	server.
		NewConfig().
		WithAppConfig().
		InitLogger().
		WithDB().
		WithRepository().
		WithWebServer().
		Start()

}