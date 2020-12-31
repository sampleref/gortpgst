package main

import (
	"flag"
	"fmt"
	gstreceiver "github.com/sampleref/gortpgst/receiver"
	lweb "github.com/sampleref/gortpgst/web"
	"net/http"
)

func init() {
	// Generate pem file for https
	lweb.GenPem()
}

func main() {

	port := flag.String("p", "9443", "https port")
	lweb.HtmlFile = flag.String("html", "", "html file absolute path")
	flag.Parse()

	// Html handle func
	http.HandleFunc("/", lweb.Web)
	go func() {
		// Support https, so we can test by lan
		fmt.Println("Web listening :" + *port)
		panic(http.ListenAndServeTLS(":"+*port, "cert.pem", "key.pem", nil))
	}()

	gstreceiver.Launch()
}
