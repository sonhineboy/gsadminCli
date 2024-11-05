package commands

import "github.com/sonhineboy/gsadminCli/tmp"

type MakeValidatorCommand struct {
	MakeBaseCommand
}

func NewMakeValidatorCommand(groupId string) *MakeValidatorCommand {
	m := new(MakeValidatorCommand)
	m.SetMakeName("validator")
	m.SetTmpFunc(tmp.ValidatorTmp)
	m.SetGroupId(groupId)

	return m
}
