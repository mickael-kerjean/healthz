package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	port := defaultString(os.Getenv("PORT"), "5341")
	fmt.Printf("starting server http://127.0.0.1:%s ...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s %s\n", time.Now().Format("20060102 15:04:05"), r.Method, r.URL, r.Header.Get("User-Agent"))
		w.Write([]byte("ok"))
	})); err != nil {
		fmt.Printf("Oops %s\n", err.Error())
	}
}

func defaultString(val string, def string) string {
	if val == "" {
		return def
	}
	return val
}
