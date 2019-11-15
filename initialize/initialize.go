package initialize

import (
	"os"

	logging "github.com/op/go-logging"
)

const commitMsgScript string = `
#!/bin/bash

komutan validate -m $2 || exit 1
`

var (
	log = logging.MustGetLogger("base")
)

type Project struct {
	homeDir string
}

func (project *Project) setProjectDir() error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	project.homeDir = workingDir
	return nil
}

func (project Project) checkCommitMsgScriptExist() error {
	gitHome := project.homeDir + "/.git/hooks"
	gitHome = gitHome
	// TODO: Control any commit_msg script exist on project git
	// TODO: If exist validate for overwrite the existing script
	return nil
}

func (project Project) writeCommitMsgScript() {

}

// Init function doing create pre commit message validation script where it is
// executed.
func Init() error {

	project := Project{}
	err := project.setProjectDir()
	if err != nil {
		log.Error("initializing failed")
		return err
	}

	return nil
}
