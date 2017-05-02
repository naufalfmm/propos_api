package main

import (
	"net/http"

	"io"

	"./models/user"
	"./models/projek"
)

type UserHandler int
type ProjekHandler int

func (u UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//io.WriteString(res,req.RequestURI)
	io.WriteString(res, user.UserController(req.RequestURI, res, req))
}

func (p ProjekHandler) ServeHTTP(res http.ResponseWriter, req, *http.Request) {
	io.WriteString(res, projek.ProjekController(req.RequestURI, res, req))
}

func main() {
	var pengg UserHandler
	var proj ProjekHandler

	mux := http.NewServeMux()
	mux.Handle("/user/", pengg)
	mux.Handle("/login/", pengg)
	mux.Handle("/projek/", proj)

	http.ListenAndServe("localhost:9000", mux)
}
