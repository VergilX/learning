package main

import (
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
