package test

import (
	"context"
	"errors"
	"log/slog"
	"os"
)

func TestFS(ctx context.Context) error {
	if os.Getenv("TEST_FS") == "no" {
		return nil
	}
	entries, err := os.ReadDir("/mnt/")
	if err != nil {
		slog.Error("fs::readdir", "error", err)
		return err
	} else if len(entries) == 0 {
		slog.Error("fs::readdir", "nothing in directory")
		return errors.New("fs::readdir unexpected length")
	}
	return nil
}
