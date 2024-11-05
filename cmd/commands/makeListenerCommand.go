package commands

import "github.com/sonhineboy/gsadminCli/tmp"

type MakeListenerCommand struct {
	MakeBaseCommand
}

func NewMakeListenerCommand(gId string) *MakeListenerCommand {
	m := new(MakeListenerCommand)
	m.SetMakeName("listener")
	m.SetTmpFunc(tmp.ListenerTmp)
	m.SetGroupId(gId)
	return m
}
