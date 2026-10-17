package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"acm.njit.edu/m/v2/backend"
)

// Arguments
var dev *bool
var port *int
var logDir *string

func setup() {
	dev = flag.Bool("dev", false, "run in development mode")
	port = flag.Int("p", 8080, "port to run server on")
	logDir = flag.String("l", "./logs", "directory to store logs")
	flag.Parse()

	if !(*dev) {
		os.MkdirAll(*logDir, os.ModePerm)
		logFilePath := fmt.Sprintf("%s/logs-%d.log", *logDir, time.Now().UnixNano())
		logFile, err := os.Create(logFilePath)
		if err != nil {
			log.Fatalf("Failed to create log: %s\n", err)
		}
		log.SetOutput(io.MultiWriter(os.Stdout, logFile))
		log.SetFlags(log.LstdFlags)
	}
}

func main() {
	setup()
	defer backend.CsvFile.Close()

	http.Handle("/", spaHandler("./dist"))
	http.HandleFunc("/api/register", backend.HandleFormSubmission)

	portStr := fmt.Sprintf(":%d", *port)

	log.Printf("[INFO] Server running on %s\n", portStr)
	log.Fatal(http.ListenAndServe(portStr, nil))
}

func spaHandler(distDir string) http.Handler {
	files := http.FileServer(http.Dir(distDir))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(path.Clean("/"+r.URL.Path), "/")
		if _, err := os.Stat(path.Join(distDir, name)); err == nil {
			files.ServeHTTP(w, r)
			return
		}

		if strings.HasPrefix(r.URL.Path, "/api/") || path.Ext(name) != "" {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, path.Join(distDir, "index.html"))
	})
}
