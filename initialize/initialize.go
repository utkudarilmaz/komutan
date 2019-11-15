package initialize

import (
	"errors"
	"os"

	logging "github.com/op/go-logging"
)

const commitMsg string = `#!/bin/bash

komutan validate -m $2 || exit 1`

var (
	log           = logging.MustGetLogger("base")
	commitMsgPath = "/.git/hooks/commit-msg"
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

func (project Project) isCommitMsgScriptExist() bool {
	filename := project.homeDir + commitMsgPath

	if _, err := os.Stat(filename); os.IsExist(err) {
		return true
	}

	return false
}

func (project Project) writeCommitMsgScript() error {
	file, err := os.Create(project.homeDir + commitMsgPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = file.Chmod(0755)
	if err != nil {
		return err
	}

	_, err = file.WriteString(commitMsg)

	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
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

	if err := project.isCommitMsgScriptExist(); err {
		return errors.New("commit-msg hook exist on initialized repository")
	}

	err = project.writeCommitMsgScript()
	if err != nil {
		return err
	}

	return nil
}
