package internal

import (
	"log"
	"net/http"
)
type Middleware func(next http.HandlerFunc) http.HandlerFunc


func ChainMiddleware(mw ...Middleware) Middleware {
    return func(final http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            last := final
            for i := len(mw) - 1; i >= 0; i-- {
                last = mw[i](last)
            }
            last(w, r)
        }
    }
}

func WithLogging(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Logged connection from %s", r.RemoteAddr)
        next.ServeHTTP(w, r)
    }
}

func WithTracing(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Tracing request for %s", r.RequestURI)
        next.ServeHTTP(w, r)
    }
}