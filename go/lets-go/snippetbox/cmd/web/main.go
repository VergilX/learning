package main

import (
    "flag"
    "log/slog"
    "net/http"
    "os"
)

// all dependencies of app
type application struct {
    logger *slog.Logger
}

func main() {
    // command line flags
    // flag function Args: 1-flagname; 2-defaultval; 3-desc
    addr := flag.String("addr", ":4000", "HTTP port address")

    // Right now addr has default value
    // Get the actual value, if present
    flag.Parse()
    
    // Structured logger
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    // application 'app' used for dependency injection
    app := &application {
        logger: logger,
    }

    // retrieve routes from routes.go into servemux
    // mux is the point of contact to routes spec
    mux := app.routes()

    logger.Info("starting server", slog.String("addr", *addr))

    err := http.ListenAndServe(*addr, mux)  // listens at port ("host:port")
    // if host not mentioned, goes an all network interfaces avaiable

    logger.Error(err.Error())
    os.Exit(1)
}
