package commands

import "github.com/sonhineboy/gsadminCli/tmp"

type MakeControllerCommand struct {
	MakeBaseCommand
}

func NewMakeControllerCommand(gId string) *MakeControllerCommand {
	m := new(MakeControllerCommand)
	m.SetMakeName("controller")
	m.SetTmpFunc(tmp.ControllerTmp)
	m.SetGroupId(gId)
	return m
}
