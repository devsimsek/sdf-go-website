package handlers

import (
	"SDF/core"
	"fmt"
	"net/http"
)

func init() {
	core.RegisterHandle("/contact", contactHandler, "GET")
}

func contactHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, core.LoadView("views/partials/head.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/navigation.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/docs/home.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/footer.html"))
	core.CheckErrorNotPanic(err)
}
