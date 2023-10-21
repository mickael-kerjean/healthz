package test

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"os"
)

func TestDNS(ctx context.Context) error {
	if os.Getenv("TEST_DNS") == "no" {
		return nil
	}
	ips, err := net.LookupIP("google.com")
	if err != nil {
		slog.Error("dns::lookup", "error", err)
		return err
	}
	if len(ips) == 0 {
		slog.Error("dns::lookup", "no matching ips")
		return errors.New("dns::lookupIP unexpected length")
	}
	return nil
}
