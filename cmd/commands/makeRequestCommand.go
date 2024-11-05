package commands

import "github.com/sonhineboy/gsadminCli/tmp"

type MakeRequestCommand struct {
	MakeBaseCommand
}

func NewMakeRequestCommand(gId string) *MakeRequestCommand {
	m := new(MakeRequestCommand)
	m.SetMakeName("request")
	m.SetTmpFunc(tmp.RequestTmp)
	m.SetGroupId(gId)
	return m
}
