package commands

import "github.com/sonhineboy/gsadminCli/tmp"

type MakeEventCommand struct {
	MakeBaseCommand
}

func NewMakeEventCommand(gId string) *MakeEventCommand {
	m := new(MakeEventCommand)
	m.SetMakeName("event")
	m.SetTmpFunc(tmp.EventTmp)
	m.SetGroupId(gId)
	return m
}
