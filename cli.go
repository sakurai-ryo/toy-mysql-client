package main

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "toy-mysql-client",
	RunE: func(cmd *cobra.Command, args []string) error {
		user, err := cmd.Flags().GetString("user")
		if err != nil {
			return err
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}

		slog.Info("args: ", "user", user, "password", password)

		if err := Client(); err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	flags := rootCmd.Flags()
	flags.StringP("user", "u", "root", "MySQL user name")
	flags.StringP("password", "p", "root", "MySQL user Password")
}
