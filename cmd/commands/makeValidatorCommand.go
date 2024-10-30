package commands

import (
	utils2 "github.com/sonhineboy/gsadminCli/pkg/utils"
	"github.com/sonhineboy/gsadminCli/tmp"
	"github.com/spf13/cobra"
	"html/template"
	"os"
	"strings"
)

var pK string

type MakeValidatorCommand struct {
	GroupId string
}

func (m *MakeValidatorCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "make:validator",
		Short:        "create a validator;",
		Example:      "make:validator [文件名]",
		SilenceUsage: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			if err := utils2.ValidateFormat(args[0]); err != nil {
				return err
			}
			if err := utils2.ValidateIsPackage(args[0], pK); err != nil {
				return err
			}
			return nil
		},
		Run: m.Run,
	}

	cmd.PersistentFlags().StringVar(&pK, "package", "", "如果在根目录创建需要传递 --package=")

	cmd.GroupID = m.GroupId

	return cmd
}

func (m *MakeValidatorCommand) Run(cmd *cobra.Command, args []string) {
	f := utils2.FileOperator{AllPath: args[0]}
	f.ParseExtAndSet()
	f.ParseName()
	f.ParseDir()
	if !f.DirExist() {
		err := f.CreateDir()
		if err != nil {
			cmd.PrintErr(err)
		}
	}

	w, err := os.Create(f.AllPath)
	if err != nil {
		cmd.PrintErr(err)
	}

	tem, err := template.New("validate").Parse(tmp.ValidatorTmp())
	if err != nil {
		cmd.PrintErr(err)
	}

	if len(pK) == 0 {
		pK = utils2.GetPackage(f.Dir)
	}

	err = tem.Execute(w, map[string]string{
		"Name":    f.NameToTitle(),
		"Package": strings.ToLower(pK),
	})
	if err != nil {
		cmd.PrintErr(err)
	}
	cmd.Printf("This is %s created", f.AllPath)
}
