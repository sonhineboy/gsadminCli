package commands

import "github.com/sonhineboy/gsadminCli/tmp"

type MakeRepositoryCommand struct {
	MakeBaseCommand
}

func NewMakeRepositoryCommand(gId string) *MakeRepositoryCommand {
	m := new(MakeRepositoryCommand)
	m.SetMakeName("repository")
	m.SetTmpFunc(tmp.RepositoryTmp)
	m.SetGroupId(gId)
	return m
}
