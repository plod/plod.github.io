package main

import (
	"github.com/elazarl/go-bindata-assetfs"
	"net/http"
	"flag"
	"log"
	"github.com/gorilla/mux"
	"os"
	"github.com/braintree/manners"
	"os/signal"
)

var port = "8080"
var addr = flag.String("addr", ":"+port, "http service address")
var hostname, _ = os.Hostname()

func addHostname(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(hostname, r.RequestURI, "was requested")
		// Set some header.
		w.Header().Add("dockwyrGwesteiwr", hostname)
		// Serve with the actual handler.
		h.ServeHTTP(w, r)
	}
}

func main() {
	mux := mux.NewRouter()
	mux.PathPrefix("/").Handler(addHostname(http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "public"})))

	go func() {
		manners.ListenAndServe(*addr, mux)
		log.Println(hostname, "listener : Started : Listening on:", *addr)
	}()


	// Listen for an interrupt signal from the OS.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println(hostname, "Shutting down...")
	log.Println(hostname, "Completed")


}