package server

import (
	"fmt"
	"time"
	"net/http"
	"testGoProject/internal"
)

type server struct {
    *http.Server
    router *http.ServeMux
    mw     internal.Middleware
}

func (s *server) Handle() {
    // handle group of routes
    internal.AddressController{
        &internal.Controller{s.router, s.mw},
    }.Handle("/api")

    //Default routes not handled by AddressController
    // What Should we do here? Error handling
    // s.router.Handle("/", s.mw(internal.FuncHandler))
}

func CreateAndListen(port string) {
    handler := http.NewServeMux()
    s := &http.Server{
        Addr:           port,
        Handler:        handler,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    mwChain := internal.ChainMiddleware(
        internal.WithLogging,
        internal.WithTracing,
    )


    ss := server{s, handler, mwChain}

    ss.Handle()
    err := ss.ListenAndServe()
    if err != nil {
        fmt.Printf("Server failed: ", err.Error())
    }

}