package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"github.com/benbjohnson/skybox/db"
	"github.com/benbjohnson/skybox/server"
)

var (
	path = flag.String("path", "", "data directory")
	port = flag.Int("port", 7000, "http port")
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: skybox [opts]")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	// Read configuration.
	flag.Usage = usage
	flag.Parse()

	// Initialize path.
	if *path == "" {
		u, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		*path = filepath.Join(u.HomeDir, ".skybox")
	}
	if err := os.MkdirAll(*path, 0700); err != nil {
		log.Fatal(err)
	}

	// Seed the PRNG for API key generation.
	rand.Seed(time.Unix64())

	// Initialize db.
	var db db.DB
	if err := db.Open(filepath.Join(*path, "db"), 0666); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start server.
	var s server.Server
	s.Addr = net.JoinHostPort("", strconv.Itoa(*port))
	s.DB = &db
	log.Printf("Listening on http://localhost%s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
