package commit

import (
	"os"
	"testing"
)

var (
	trueMessages = []string{
		"feat(build): new feature",
		"fix: -word counte!r",
		"Merge branch 'prod-imp-master-and-restaurant-panel-monetary-operation-",
		"refactor(build): test3 . wEqW_w",
		"feat(build-test)!: so2mething",
		"fix!: word count''er2",
		"feat(build): .new feature",
		`refactor: commit msg validator name changed and validation steps added

		ValidateCommitMsgString function name change to ValidateCommitMsg. Also validation process splitting some steps. Commit message's header section character limit controlling after now
`,
	}

	falseMessages = []string{
		"fix: Word counter",
		"refactor(build-):",
		"merge branchwewe",
		"feat(build-test): something.",
		"some more thing",
		"crazy: new feature",
		"fix more powWWW3213er",
		"fix: more *powWWW3213er!",
		`refactor: commit msg validator name changed and validation steps added1322333333333333333333333333333333333

		ValidateCommitMsgString function name change to ValidateCommitMsg. Also validation process splitting some steps. Commit message's header section character limit controlling after now`,
	}
)

func TestValidateCommitMsg(t *testing.T) {

	for index := 0; index < len(trueMessages); index++ {
		err := ValidateCommitMsg(trueMessages[index])
		if err != nil {
			t.Errorf(
				"%s message has true format but can't pass the validation",
				trueMessages[index],
			)
		}
	}

	for index := 0; index < len(falseMessages); index++ {
		err := ValidateCommitMsg(falseMessages[index])
		if err == nil {
			t.Errorf(
				"%s message has wrong template but pass the validation",
				falseMessages[index],
			)
		}
	}
}

func TestValidateCommitMsgFromFile(t *testing.T) {
	file, err := os.Create("/tmp/commit")
	if err != nil {
		t.Errorf(
			"error occurred when creating file: %s",
			err.Error(),
		)
	}
	defer file.Close()

	err = file.Chmod(0755)
	if err != nil {
		t.Errorf(
			"error occurred when changing file permission bits: %s",
			err.Error(),
		)
	}

	_, err = file.WriteString(trueMessages[5])
	if err != nil {
		t.Errorf(
			"error occurred when writing commit message to file: %s",
			err.Error(),
		)
	}

	err = file.Sync()
	if err != nil {
		t.Errorf(
			"error occurred when committing to file: %s",
			err.Error(),
		)
	}

	err = ValidateCommitMsgFromFile("/tmp/commit")
	if err != nil {
		t.Errorf(
			"%s message readed from a file but some error occurred: %s",
			trueMessages[5],
			err.Error(),
		)
	}
}
