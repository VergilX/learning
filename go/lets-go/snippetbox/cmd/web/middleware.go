package main

import (
    "fmt"
    "net/http"
)

func commonHeaders(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        // set header values
        w.Header().Set("Content-Security-Policy",
        "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")

        w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "deny")
        w.Header().Set("X-XSS-Protection", "0")
        w.Header().Set("Server", "Go")

        // call the ServeHTTP func of next handler func
        next.ServeHTTP(w, r)
    }

    return http.HandlerFunc(fn)
}

func (app *application) logRequest(next http.Handler) http.Handler {
    // shorthand similar functionality of above func
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var (
            ip      = r.RemoteAddr
            proto   = r.Proto
            method  = r.Method
            uri     = r.URL.RequestURI()
        )

        app.logger.Info("received request", "ip", ip, "proto", proto, "method", method, "uri", uri)

        next.ServeHTTP(w, r)
    })
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // deferred functions are always called during panic
        // unwinding. So use it for elegant exit
        defer func() {
            // builtin recover() detects panic
            if err := recover(); err != nil {
                w.Header().Set("Connection", "close")

                app.serverError(w, r, fmt.Errorf("%s", err))
            }
        }()

        next.ServeHTTP(w, r)
    })
}
