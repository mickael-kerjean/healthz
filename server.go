package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/mickael-kerjean/healthz/test"
)

func main() {
	slog.Default()
	port := defaultString(os.Getenv("PORT"), "5341")
	fmt.Printf("starting server http://127.0.0.1:%s ...\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), http.HandlerFunc(handler)); err != nil {
		fmt.Printf("Oops %s\n", err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	var errs []error
	if err := test.TestDial(r.Context()); err != nil {
		errs = append(errs, err)
	}
	if err := test.TestDNS(r.Context()); err != nil {
		errs = append(errs, err)
	}
	if err := test.TestK8s(r.Context()); err != nil {
		errs = append(errs, err)
	}
	w.Header().Add("Content-Type", "text/plain")
	if len(errs) == 0 {
		slog.Info("HTTP", "status", 200, "time", time.Since(now)/1000/1000, "method", r.Method, "url", r.URL, "user-agent", r.Header.Get("User-Agent"))
		w.Write([]byte("OK"))
	} else {
		slog.Error("HTTP", "status", 500, "time", time.Since(now), "method", r.Method, "url", r.URL, "user-agent", r.Header.Get("User-Agent"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR"))
	}
}

func defaultString(val string, def string) string {
	if val == "" {
		return def
	}
	return val
}
