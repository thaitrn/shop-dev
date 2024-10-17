package main

import initalize "shop-dev/internal/0.initalize"

func main() {
	// r := routers.NewRouter()

	// r.Run(":8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	initalize.Run()
}
