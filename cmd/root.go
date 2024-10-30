package cmd

import (
	"github.com/sonhineboy/gsadminCli/cmd/commands"
	"github.com/spf13/cobra"
)

const Version = "v0.0.1"

func Execute() {
	rootCmd := &cobra.Command{
		Use:     "gsadminCli",
		Long:    "gsadminCli is primarily a scaffolding tool for gsadmin, which can help speed up your development process.",
		Version: Version,
	}

	rootCmd.AddGroup(
		&cobra.Group{
			ID:    Validator.ID,
			Title: Validator.ID,
		},
	)

	command := &commands.MakeValidatorCommand{GroupId: Validator.ID}
	rootCmd.AddCommand(
		command.Command(),
	)
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}
