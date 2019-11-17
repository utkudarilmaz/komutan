package commit

import (
	"os"
	"testing"
)

var (
	trueMessages = []string{
		"feat(build): new feature",
		"fix: -word counter",
		"refactor(build): test3",
		"feat(build-test): so2mething",
	}

	falseMessages = []string{
		"feat(build): .new feature",
		"fix: Word counter",
		"refactor(build-):",
		"feat(build-test): something.",
		"some more thing",
		"crazy: new feature",
		"fix more powWWW3213er",
		"fix: more *powWWW3213er!",
	}
)

func TestValidateCommitMsgString(t *testing.T) {

	for index := 0; index < len(trueMessages); index++ {
		err := ValidateCommitMsgString(trueMessages[index])
		if err != nil {
			t.Errorf(
				"%s message has true format but can't pass the validation",
				trueMessages[index],
			)
		}
	}

	for index := 0; index < len(falseMessages); index++ {
		err := ValidateCommitMsgString(falseMessages[index])
		if err == nil {
			t.Errorf(
				"%s message has wrong template but pass the validation",
				falseMessages[index],
			)
		}
	}
}

func TestValidateCommitMsgFile(t *testing.T) {
	file, err := os.Create("/tmp/commit")
	if err != nil {
		t.Errorf(
			"error occured when creating file: %s",
			err.Error(),
		)
	}
	defer file.Close()

	err = file.Chmod(0755)
	if err != nil {
		t.Errorf(
			"error occured when changing file permission bits: %s",
			err.Error(),
		)
	}

	_, err = file.WriteString(trueMessages[0])
	if err != nil {
		t.Errorf(
			"error occured when writing commit message to file: %s",
			err.Error(),
		)
	}

	err = file.Sync()
	if err != nil {
		t.Errorf(
			"error occured when commiting to file: %s",
			err.Error(),
		)
	}

	err = ValidateCommitMsgFile("/tmp/commit")
	if err != nil {
		t.Errorf(
			"%s message readed from a file but some error occured: %s",
			trueMessages[0],
			err.Error(),
		)
	}
}
