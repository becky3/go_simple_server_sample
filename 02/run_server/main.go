package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))

	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	fmt.Fprintf(w, "<html><body>\n")

	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<p>二回目以降</p>\n")
	} else {
		fmt.Fprintf(w, "<p>初訪問</p>")
	}

	fmt.Fprintf(w, "</body></html>\n")

}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
