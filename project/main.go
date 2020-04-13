package main

import (
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	// ip addr, server mux(muliplexer)
	http.ListenAndServe(":3000", nil)
}
