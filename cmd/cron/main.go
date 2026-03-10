package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/robfig/cron/v3"
)

func main() {
	// Parse flags first
	runNow := flag.Bool("run-now", false, "Run job immediately on startup")
	flag.Parse()

	// After flag.Parse(), non-flag arguments are in flag.Args()
	args := flag.Args()

	// We need at least 2 arguments: <schedule> and at least 1 part of a <command>
	if len(args) < 2 {
		slog.Error("usage: ./cron [--run-now] <schedule> <command>")
		os.Exit(1)
	}

	schedule := args[0]
	// Join all remaining arguments into a single command string
	// This allows: ls -al -> "ls -al"
	command := strings.Join(args[1:], " ")

	slog.Info("Config",
		"schedule", schedule,
		"command", command,
		"run_now", *runNow,
	)

	// Initialize Cron with Seconds support and "Skip If Still Running" protection
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.SkipIfStillRunning(cron.DefaultLogger),
		),
	)

	// Define the job execution logic
	jobFunc := func() {
		// Using /bin/sh -c allows for complex commands, pipes, and redirects
		cmd := exec.Command("/bin/sh", "-c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			slog.Error("command failed", "error", err)
			return
		}
	}

	// Register the job to the scheduler
	_, err := c.AddFunc(schedule, jobFunc)
	if err != nil {
		slog.Error("failed to add cron job", "error", err)
		os.Exit(1)
	}

	c.Start()

	// Execute immediately if requested
	if *runNow {
		go jobFunc()
	}

	// Wait for termination signals (Ctrl+C or SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	slog.Info("Shutdown signal received")

	// Stop the scheduler gracefully
	c.Stop()
}
