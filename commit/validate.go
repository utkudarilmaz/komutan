package commit

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"

	logging "github.com/op/go-logging"
)

var (
	log             = logging.MustGetLogger("base")
	defaultTemplate = `^(feat|fix|refactor|chore)(\([a-zA-Z0-9]*-?[a-zA-z0-9]+\))?:\s[a-z].([a-zA-Z0-9\.',_-]|\s)+[^\.\!\?=_-]$`
)

// ValidateCommitMsgFile is validate the commit message where is given file
func ValidateCommitMsgFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if matched, err := regexp.Match(defaultTemplate, buffer); !matched {
		if err != nil {
			return err
		}
		return errors.New("\"" + string(buffer[:len(buffer)-1]) + "\"" + " commit message is not valid")
	}

	log.Notice("\"" + string(buffer[:len(buffer)-1]) + "\"" + " commit message is valid")

	return nil
}

// ValidateCommitMsgString is validating the given commit message with following RegExp.
// ^(feat|fix|refactor|chore)(\([a-zA-Z0-9]*-?[a-zA-z0-9]+\))?:\s[a-z].([a-zA-Z0-9\.,_-]|\s)+[^\.\!\?=_-]$
// RegExp rules produced based a couple policies where the policies defined
// https://www.conventionalcommits.org
func ValidateCommitMsgString(message string) error {
	if len(message) > 72 || len(message) < 7 {
		return errors.New("message length must between 7 and 72 character")
	}

	matched, _ := regexp.MatchString(defaultTemplate, message)
	if !matched {
		return errors.New("\"" + message + "\"" + " commit message is not valid")
	}

	log.Notice("\"" + message + "\"" + " commit message is valid")
	return nil
}
