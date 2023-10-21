package test

import (
	"context"
	"log/slog"
	"net"
	"os"
)

func TestDial(ctx context.Context) error {
	if os.Getenv("TEST_DIAL") == "no" {
		return nil
	}
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		slog.Error("dial::udp 8.8.8.8:53", "error", err)
		return err
	}
	return conn.Close()
}
