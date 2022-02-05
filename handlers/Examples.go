package handlers

import (
	"SDF/core"
	"fmt"
	"net/http"
)

func init() {
	core.RegisterHandle("/examples", examplesHome, "GET")
	// core.RegisterHandle("/e/", examplesHandler, "GET") soon...
}

func examplesHome(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, core.LoadView("views/partials/head.html", core.View{Page: "Examples"}))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/navigation.html", core.View{Slug: "examples"}))
	_, err = fmt.Fprintf(writer, core.LoadView("views/examples/home.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/footer.html"))
	core.CheckErrorNotPanic(err)
}

/** Soon...
func examplesHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, core.LoadView("views/partials/head.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/navigation.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/docs/home.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/footer.html"))
	core.CheckErrorNotPanic(err)
}
*/
