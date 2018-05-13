package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
	"github.com/zserge/webview"
)

var webviewSettings = webview.Settings{
	Title:                  "Desktop by Machine Box",
	Width:                  650,
	Height:                 350,
	Resizable:              true,
	ExternalInvokeCallback: handleClientInvoke,
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

// handleClientInvoke handles client-side calls to
// window.external.invoke.
func handleClientInvoke(w webview.WebView, data string) {
	webview.Debug("TODO: handleClientInvoke: ", data)
}

func run() error {
	const endpoint = "0.0.0.0:0"
	debug := os.Getenv("MB_DESKTOP_DEBUG") == "true"
	box := packr.NewBox("./www")
	http.HandleFunc("/time", handleTime)
	http.Handle("/", http.FileServer(box))
	ln, err := net.Listen("tcp", endpoint)
	if err != nil {
		return errors.Wrap(err, "net.Listen")
	}
	go func() {
		if err := http.Serve(ln, nil); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(2)
		}
	}()
	addr := "http://" + ln.Addr().String()
	log.Println("serving via", addr)
	client := http.Client{Timeout: 100 * time.Millisecond}
	for {
		time.Sleep(100 * time.Millisecond)
		resp, err := client.Get(addr)
		if err != nil {
			continue
		}
		resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			break
		}
	}
	webviewSettings.URL = addr
	webviewSettings.Debug = debug
	win := webview.New(webviewSettings)
	win.Run()
	return nil
}

// handleTime handles the /time endpoint.
// It is used to demonstrate an API call from app.js.
func handleTime(w http.ResponseWriter, r *http.Request) {
	if _, err := io.WriteString(w, time.Now().Format(time.RFC822)); err != nil {
		log.Println(err)
	}
}
