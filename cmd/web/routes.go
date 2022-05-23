package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	standardMiddleWare := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleWare := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()
	mux.Get("/", dynamicMiddleWare.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleWare.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippetForm))
	mux.Get("/snippet/:id", dynamicMiddleWare.ThenFunc(app.showSnippet))
	mux.Get("/user/signup", dynamicMiddleWare.ThenFunc(app.signupUserForm))
	mux.Get("/user/login", dynamicMiddleWare.ThenFunc(app.loginUserForm))

	mux.Post("/snippet/create", dynamicMiddleWare.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippet))
	mux.Post("/user/signup", dynamicMiddleWare.ThenFunc(app.signupUser))
	mux.Post("/user/login", dynamicMiddleWare.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleWare.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleWare.Then(mux)
}
