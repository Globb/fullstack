package cmd

import "github.com/spf13/cobra"

func newRootCmd() *cobra.Command {
	var debug bool
	var domain string
	var port int

	rootCmd := &cobra.Command{
		Use:   "backend",
		Short: "backend cli",
		Long:  "backend cli",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.PersistentFlags().BoolVar(
		&debug,
		"debug",
		false,
		"enable debug logging",
	)

	rootCmd.PersistentFlags().StringVar(
		&domain,
		"domain",
		"http://localhost",
		"Domain url for server.",
	)

	rootCmd.PersistentFlags().IntVar(
		&port,
		"port",
		8080,
		"Port used to connect to domain url for server.",
	)

	rootCmd.AddCommand(newBackendClientCmd(&domain, &port, &debug))
	rootCmd.AddCommand(newBackendServerCmd(&domain, &port, &debug))

	return rootCmd
}

func Execute() error {
	err := newRootCmd().Execute()
	if err != nil {
		return err
	}

	return nil
}

