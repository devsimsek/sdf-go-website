package handlers

import (
	"SDF/core"
	docs "SDF/core/libraries/database"
	"fmt"
	"net/http"
	"strings"
)

func init() {
	core.RegisterStandaloneHandle("/search", searchHandle)
	// For Testing :D core.RegisterHandle("/testHandle", testHandle, "GET")
	docs.Open("docs.edb", "databases")
}

func searchHandle(writer http.ResponseWriter, request *http.Request) {

	// Search
	keys, ok := request.URL.Query()["sq"]
	if !ok || len(keys[0]) < 1 {
		http.Redirect(writer, request, "/?error=0x", http.StatusSeeOther)
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	query := docs.Read()
	var q map[string]interface{}
	for s, i := range query {
		if i.(map[string]interface{})["SubSlug"] == strings.ToLower(keys[0]) {
			q = map[string]interface{}{
				"title":   i.(map[string]interface{})["title"],
				"slug":    i.(map[string]interface{})["slug"],
				"SubSlug": i.(map[string]interface{})["SubSlug"],
				"route":   s,
			}
			break
		} else {
			q = map[string]interface{}{
				"title":    "Not Found",
				"slug":     "not_found",
				"SubSlug":  "not_found",
				"pageData": "<div class=\"alert alert-danger\" role=\"alert\">\n    <div class=\"d-flex align-items-center\">\n        <i class=\"bi bi-exclamation-circle-fill fs-1\"></i>\n        <h1 class=\"alert-heading fw-bolder d-inline\">&nbsp;Error, Query Not Found.</h1>\n    </div>\n    <hr>\n    <p class=\"mb-0 fs-5\" id=\"message\">\n        Query That You've Searched For Does Not Exists In Sdf Documents Database.\n    </p>\n</div>\n</main>",
			}
		}
	}

	_, err := fmt.Fprintf(writer, core.LoadView("views/partials/head.html", core.View{Page: "Docs"}))
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/navigation.html", core.View{Slug: "docs"}))
	_, err = fmt.Fprintf(writer, core.LoadView("views/docs/partials/sidebar.html", core.View{Slug: "home", SubSlug: "home"}))
	if q["slug"].(string) == "not_found" {
		_, err = fmt.Fprintf(writer, q["pageData"].(string))
	} else {
		_, err := fmt.Fprintf(writer, "<html><head><meta http-equiv='refresh' content='0; url="+"/d/"+q["route"].(string)+"'><title>Redirecting...</title></head><body>Redirecting...<script>window.location.replace('"+"/d/"+q["route"].(string)+"')</script></body></html>")
		core.CheckErrorNotPanic(err)
	}
	_, err = fmt.Fprintf(writer, core.LoadView("views/partials/footer.html"))
	core.CheckErrorNotPanic(err)
}
