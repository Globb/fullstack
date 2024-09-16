package cmd

import (
	"fullstack/backend/api"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

func newBackendServerCmd(domain *string, port *int, debug *bool) *cobra.Command {
	backendServerCmd := &cobra.Command{
		Use:   "server",
		Short: "backend server",
		Long:  "backend server",
		Run: func(cmd *cobra.Command, args []string) {
			runServer(domain, port, debug)
		},
	}

	return backendServerCmd
}

func runServer(domain *string, port *int, debug *bool) {
	//log level default is info
	loggerOptions := &slog.HandlerOptions{}
	if *debug {
		loggerOptions = &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
	}

	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, loggerOptions),
	)
	slog.SetDefault(logger)

	server, err := api.NewServer(*domain, *port, logger)
	if err != nil {
		logger.Error("failed to create server",
			slog.String("error", err.Error()),
		)
		return
	}
	logger.Debug("server interface created")

	server.RunServer()
}
