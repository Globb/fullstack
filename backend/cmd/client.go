package cmd

import (
	"fullstack/backend/api"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

func newBackendClientCmd(domain *string, port *int, debug *bool) *cobra.Command {
	var file string
	var jsonString string
	backendClientCmd := &cobra.Command{
		Use:   "client",
		Short: "backend client",
		Long:  "backend client",
		Run: func(cmd *cobra.Command, args []string) {
			runClient(file, jsonString, domain, port, debug)
		},
	}

	backendClientCmd.PersistentFlags().StringVar(
		&file,
		"file",
		"",
		"Provide input as a json file. If used with json flag, this input file will be ignored.",
	)

	backendClientCmd.PersistentFlags().StringVar(
		&jsonString,
		"json",
		"",
		"Provide input as raw json. NOTE: you might need to escape with \\\"",
	)

	return backendClientCmd
}

func runClient(
	file string,
	jsonString string,
	domain *string,
	port *int,
	debug *bool,
) {
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

	client, err := api.NewClient(*domain, *port, logger)
	if err != nil {
		logger.Error("failed to create client",
			slog.String("error", err.Error()),
		)
		return
	}

	logger.Debug("client interface created")
	if jsonString != "" {
		logger.Debug("using json flag for input")
		input, err := api.ValidateInputRawJson(jsonString)
		if err != nil {
			logger.Error(
				"input failed to validate, all fields must be valid to submit input",
				slog.String("validateType", "raw json"),
				slog.Int("AdCampaignId", input.AdCampaignId),
				slog.Int("CustomerId", input.CustomerId),
				slog.String("GameName", input.GameName),
				slog.String("ImageName", input.ImageName),
				slog.Bool("ValidAccount", input.ValidAccount),
				slog.String("error", err.Error()),
			)
			return
		}

		client.SubmitInput(input)
		return
	}

	logger.Debug("using file flag for input")
	input, err := api.ValidateInputFile(file)
	if err != nil {
		logger.Error(
			"input failed to validate, all fields must be valid to submit input",
			slog.String("validateType", "file"),
			slog.Int("AdCampaignId", input.AdCampaignId),
			slog.Int("CustomerId", input.CustomerId),
			slog.String("GameName", input.GameName),
			slog.String("ImageName", input.ImageName),
			slog.Bool("ValidAccount", input.ValidAccount),
			slog.String("error", err.Error()),
		)
		return
	}
	client.SubmitInput(input)
}
