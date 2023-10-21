package test

import (
	"context"
	"crypto/tls"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
)

func TestK8s(ctx context.Context) error {
	if os.Getenv("TEST_K8S") == "no" {
		return nil
	}

	// detect kube-dns issue
	ips, err := net.LookupIP("kubernetes.default.svc.cluster.local")
	if err != nil {
		slog.Error("k8s::lookup kubernetes.default.svc.cluster.local", "err", err)
		return err
	}
	if len(ips) == 0 {
		slog.Error("k8s::lookup kubernetes.default.svc.cluster.local", "len", 0)
		return errors.New("dns::lookupIP unexpected length")
	}

	// coreDNS potential issues to be logged
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	resp, err := client.Get("https://kubernetes.default.svc.cluster.local:443/healthz")
	if err != nil {
		slog.Warn("k8s::healthz kubernetes.default.svc.cluster.local", "err", err)
		return nil
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		slog.Warn("k8s::healthz kubernetes.default.svc.cluster.local", "status", resp.StatusCode)
		return nil
	}
	return nil
}
