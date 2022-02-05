package handlers

import (
	"SDF/core"
	docs "SDF/core/libraries/database"
	"fmt"
	"net/http"
)

func init() {
	core.RegisterStandaloneHandle("/d/", docHandler)
	core.RegisterHandle("/docs", docHome, "GET")
	// For Testing :D core.RegisterHandle("/testHandle", testHandle, "GET")
	docs.Open("docs.edb", "databases")
}

func docHome(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, core.LoadView("views/partials/head.html", core.View{Page: "Docs"}))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/navigation.html", core.View{Slug: "docs"}))
	_, err = fmt.Fprintf(writer, core.LoadView("views/docs/partials/sidebar.html", core.View{Slug: "home", SubSlug: "home"}))
	_, err = fmt.Fprintf(writer, core.LoadView("views/docs/home.html"))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/footer.html"))
	core.CheckErrorNotPanic(err)
}

func docHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path[3:] != "" {
		// Compile Request
		topic := docs.Get(request.URL.Path[3:])

		if topic != nil {
			_, err := fmt.Fprintf(writer, core.LoadView("views/partials/head.html", core.View{Page: "Docs | " + topic.(map[string]interface{})["title"].(string)}))
			_, err = fmt.Fprintf(writer, core.LoadView("views/partials/navigation.html", core.View{Slug: "docs"}))
			_, err = fmt.Fprintf(writer, core.LoadView("views/docs/partials/sidebar.html", core.View{
				Slug:    topic.(map[string]interface{})["slug"].(string),
				SubSlug: topic.(map[string]interface{})["SubSlug"].(string),
			}))
			_, err = fmt.Fprintf(writer, topic.(map[string]interface{})["pageData"].(string))
			_, err = fmt.Fprintf(writer, core.LoadView("views/partials/footer.html"))
			core.CheckErrorNotPanic(err)
		} else {
			_, err := fmt.Fprintf(writer, "<html><head><meta http-equiv='refresh' content='0; url=/?error=0xednf404'><title>Redirecting...</title></head><body>Redirecting...<script>window.location.replace('/?error=0xednf404')</script></body></html>")
			core.CheckErrorNotPanic(err)
			// Document Not Found Error Code;
			http.Redirect(writer, request, "/?error=0xednf404", http.StatusSeeOther)
		}
	}
}

/**
For Testing;
func testHandle(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, core.LoadView("views/partials/head.html", core.View{Page: "Docs | Home"}))
	_, err = fmt.Fprintf(w, core.LoadView("views/partials/navigation.html", core.View{Slug: "docs"}))
	_, err = fmt.Fprintf(w, core.LoadView("views/docs/partials/sidebar.html", core.View{Slug: "home"}))
	_, err = fmt.Fprintf(w, core.LoadView("views/docs/view.html"))
	_, err = fmt.Fprintf(w, core.LoadView("views/partials/footer.html"))
	core.CheckErrorNotPanic(err)
}
*/
