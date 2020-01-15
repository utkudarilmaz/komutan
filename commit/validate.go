package commit

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	logging "github.com/op/go-logging"
)

var (
	log = logging.MustGetLogger("base")
)

// ValidateCommitMsgFromFile is validate the commit message where is given file
func ValidateCommitMsgFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = ValidateCommitMsg(string(buffer[:]))
	if err != nil {
		return err
	}

	return nil
}

func validateType(message string) error {
	typeTemplate := regexp.MustCompile(`^(feat|docs|style|perf|test|fix|refactor|chore){1}(\([a-zA-Z0-9]+(-?)[a-zA-z0-9]+\))?!?:\ [a-z-.]{1}.+[^\.,\s!\?\\\ \{\}\[\]]+$`)
	var lines = strings.Split(message, "\n")

	index := typeTemplate.FindStringIndex(lines[0])
	if index == nil {
		return errors.New("\"" + message + "\"" + " commit message is not valid")
	} else if index[0] != 0 {
		return errors.New("\"" + message + "\"" + " commit message's type or optional scope is not valid")
	} else if index[1] > 72 {
		return errors.New("\"" + message + "\"" + " commit message can't be more than 72 character")
	}

	return nil
}

func validateMerge(message string) error {
	mergeTemplate := regexp.MustCompile(`^(Merge\ branch\ ).+$`)
	var lines = strings.Split(message, "\n")

	index := mergeTemplate.FindStringIndex(lines[0])
	if index == nil {
		return errors.New("\"" + message + "\"" + " commit message is not valid")
	}

	return nil
}

// ValidateCommitMsg is validating the given commit message looking
// some regexp rules.
// RegExp rules produced based a couple policies where the policies defined
// https://www.conventionalcommits.org
func ValidateCommitMsg(message string) error {

	err := validateType(message)
	if err != nil {
		errX := validateMerge(message)
		if errX != nil {
			return err
		}
	}

	log.Notice("\"" + message + "\"" + " commit message is valid")
	return nil
}
