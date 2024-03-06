package main

import "net/http"

func (app *application) routes() http.Handler {
	// router - controller
	mux := http.NewServeMux()

	filseServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", filseServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	mux.HandleFunc("/download", app.downloadLogo)

	return app.logRequest(secureHeaders(mux))
}
