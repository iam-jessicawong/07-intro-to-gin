package main

import "07-gin-get-started/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
