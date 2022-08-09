package main

import "icl2-external-server-sample/internal/api"

func main() {
	api.StartServer()

	select {} // pending...
}
