package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/skybox/skybox/db"
	"github.com/skybox/skybox/server"
)

var (
	addr     = flag.String("addr", ":80", "HTTP address")
	tlsAddr  = flag.String("tls-addr", ":443", "HTTPS address")
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
	var useTLS = (*certFile != "" && *keyFile != "")

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

	// Initialize handler.
	h, err := server.NewHandler(&db)
	if err != nil {
		log.Fatal(err)
	}

	// Start servers.
	log.Printf("Listening on http://localhost%s", *addr)
	if useTLS {
		log.Printf("Listening on http://localhost%s", *tlsAddr)
	}
	log.SetFlags(log.LstdFlags)

	go func() { log.Fatal(http.ListenAndServe(*addr, h)) }()
	if useTLS {
		go func() { log.Fatal(http.ListenAndServeTLS(*tlsAddr, *certFile, *keyFile, h)) }()
	}

	select {}
}
