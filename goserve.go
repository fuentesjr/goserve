package main

import (
  "log"
  "net/http"
  "flag"
)

func main() {
	var dpath = flag.String("dpath", ".", "path to directory that holds static assets")
	var port = flag.String("port", "8080", "http port to listen on")
	flag.Parse()

  fs := http.FileServer(http.Dir(*dpath))
  http.Handle("/", fs)

  log.Printf("Serving files from: %s\n", *dpath)
  log.Printf("Listening on port %s ...\n", *port)
  http.ListenAndServe(":" + *port, nil)
}
