package handlers

import (
	"SDF/core"
	"fmt"
	"net/http"
)

func init() {
	core.RegisterHandle("/home", homeHandler, "GET")
	core.RegisterHandle("/", homeHandler, "GET")
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, core.LoadView("views/partials/head.html", core.View{
		Page: "Home",
	}))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/navigation.html", core.View{
		Slug: "home",
	}))
	_, err = fmt.Fprintf(writer, core.LoadView("views/home.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/footer.html"))
	core.CheckErrorNotPanic(err)
}
