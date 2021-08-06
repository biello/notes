// Usage
//
//   wiki -db="/tmp/foo.db" -port=8080
//
package main

import (
	"flag"
	"time"

	"github.com/biello/notes/db"
	"github.com/biello/notes/helper"
	"github.com/biello/notes/server"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var dbFile string
var port string
var loglevel string

func main() {
	flag.StringVar(&dbFile, "db", "/tmp/notes.db", "Path to the BoltDB file")
	flag.StringVar(&port, "port", "8080", "Http server listen port")
	flag.StringVar(&loglevel, "log", "debug", "log level")
	flag.Parse()

	// Setup the logger used by the server
	logger := helper.NewLogger(loglevel)

	// Setup the database used by the server
	db, err := newDB(dbFile)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	server.Init(logger, db)
	server.HTTP(r)

	appServer := endless.NewServer(":"+port, r)
	appServer.ReadTimeout = 10 * time.Second
	appServer.WriteTimeout = 10 * time.Second
	appServer.IdleTimeout = 2 * time.Minute
	logger.Println("Listening on http://0.0.0.0:" + port)
	if err := appServer.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func newDB(fn string) (*db.DB, error) {
	db := &db.DB{}

	if err := db.Open(dbFile, 0600); err != nil {
		return nil, err
	}

	return db, nil
}
