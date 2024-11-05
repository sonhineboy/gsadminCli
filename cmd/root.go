package cmd

import (
	"github.com/sonhineboy/gsadminCli/cmd/commands"
	"github.com/spf13/cobra"
)

const Version = "v0.1.6"

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

	validator := commands.NewMakeValidatorCommand(Validator.ID)
	request := commands.NewMakeRequestCommand(Validator.ID)
	event := commands.NewMakeEventCommand(Validator.ID)
	listener := commands.NewMakeListenerCommand(Validator.ID)
	controller := commands.NewMakeControllerCommand(Validator.ID)
	repository := commands.NewMakeRepositoryCommand(Validator.ID)
	model := commands.NewMakeModelCommand(Validator.ID)

	rootCmd.AddCommand(
		validator.Command(),
		request.Command(),
		event.Command(),
		listener.Command(),
		controller.Command(),
		repository.Command(),
		model.Command(),
	)

	err := rootCmd.Execute()
	if err != nil {
		return
	}
}
