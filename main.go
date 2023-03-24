package main

import "challange-06/routers"

func main() {
	var PORT = ":8090"
	routers.StartServer().Run(PORT)
}
