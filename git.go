package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {
	// JSON output (for production)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	for {

		select {
		case <-time.Tick(time.Second):

			logger.Info("User logged in", "user_id", 123, "ip", "192.168.1.1")
		}
	}

}
