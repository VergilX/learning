package main

import (
    "database/sql"
    "flag"
    "log/slog"
    "net/http"
    "os"

    "github.com/VergilX/learning/go/lets-go/snippetbox/internal/models"

    _"github.com/go-sql-driver/mysql"
)

// all dependencies of app (dependency injection)
type application struct {
    logger *slog.Logger
    snippets *models.SnippetModel
}

func main() {
    // command line flags
    // flag function Args: 1-flagname; 2-defaultval; 3-desc
    addr := flag.String("addr", ":4000", "HTTP port address")
    dsn := flag.String("dsn", "web:@/snippetbox?parseTime=true", "MySQL data source name")  // cmd for DB DSN String (Data Source Name)

    // Right now cmd has default values
    // Get the actual value, if present
    flag.Parse()
    
    // Structured logger
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    // Connection pool to database
    db, err := openDB(*dsn)
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }

    // defer close (delay until main is done)
    defer db.Close()

    // application 'app' used for dependency injection
    app := &application {
        logger:     logger,
        snippets:   &models.SnippetModel{db: DB},
    }

    // retrieve routes from routes.go into servemux
    // mux is the point of contact to routes spec
    mux := app.routes()

    logger.Info("starting server", slog.String("addr", *addr))

    err = http.ListenAndServe(*addr, mux)  // listens at port ("host:port")
    // if host not mentioned, goes an all network interfaces avaiable

    logger.Error(err.Error())
    os.Exit(1)
}


func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}
