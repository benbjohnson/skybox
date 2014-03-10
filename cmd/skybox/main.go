package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/skybox/skybox/db"
	"github.com/skybox/skybox/server"
)

var (
	dataDir  = flag.String("data-dir", "", "data directory")
	certFile = flag.String("cert-file", "", "SSL certificate file")
	keyFile  = flag.String("key-file", "", "SSL key file")
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: skybox [opts]")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	log.SetFlags(0)

	// Read configuration.
	flag.Usage = usage
	flag.Parse()

	// Validate arguments.
	if *dataDir == "" {
		log.Fatal("data directory required: --data-dir")
	} else if *certFile != "" && *keyFile == "" {
		log.Fatal("key file required: --key-file")
	} else if *keyFile != "" && *certFile == "" {
		log.Fatal("certificate file required: --cert-file")
	}

	// Initialize data directory.
	if err := os.MkdirAll(*dataDir, 0700); err != nil {
		log.Fatal(err)
	}

	// Seed the PRNG for API key generation.
	rand.Seed(time.Now().Unix())

	// Initialize db.
	var db db.DB
	if err := db.Open(filepath.Join(*dataDir, "db"), 0666); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start server.
	var s server.Server
	s.DB = &db
	log.Printf("Listening on http://localhost%s", s.Addr)
	log.SetFlags(log.LstdFlags)
	log.Fatal(s.ListenAndServe())
}
