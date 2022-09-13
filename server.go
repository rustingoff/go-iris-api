package main

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"time"
)

func NewHTTPServer(app *iris.Application) *http.Server {
	return &http.Server{
		Addr:           ":8080",
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
