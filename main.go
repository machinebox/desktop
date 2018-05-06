package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zserge/webview"
)

//go:generate go-bindata-assetfs www/...

func main() {
	go func() {
		http.Handle("/", http.FileServer(assetFS()))
		if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(2)
		}
	}()
	if err := webview.Open("Desktop by Machine Box", "http://0.0.0.0:8080", 1024, 768, true); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}
