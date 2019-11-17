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

