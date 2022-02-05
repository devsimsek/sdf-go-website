package handlers

import (
	"SDF/core"
	"fmt"
	"net/http"
)

func init() {
	core.RegisterHandle("/p/", pageHandler, "GET")
}

func pageHandler(writer http.ResponseWriter, request *http.Request) {

	// parse page by request.URL.Path[1:]

	_, err := fmt.Fprintf(writer, core.LoadView("views/partials/head.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/navigation.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/docs/home.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/footer.html"))
	core.CheckErrorNotPanic(err)
}
