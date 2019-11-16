package commit

import (
	"errors"
	"regexp"

	logging "github.com/op/go-logging"
)

var (
	log = logging.MustGetLogger("base")
	defaultTemplate = `^(feat|fix|refactor|chore)(\([a-zA-Z0-9]*-?[a-zA-z0-9]+\))?:\s[a-z].([a-zA-Z0-9\.',_-]|\s)+[^\s\.\!\?=_-]$`
)

// Validate the given commit message with following RegExp.
// ^(?=.{1,72}$)(feat|fix|refactor|chore)(|\([a-zA-Z0-9]+-?[a-zA-z0-9]+)\):\s[a-z].([a-zA-Z0-9\.,_-]|\s)+[^\.\!\?=_-]$
// RegExp rules produced based a couple policies where the policies defined
// https://www.conventionalcommits.org
func Validate(message string) error {
	if len(message) > 72 || len(message) < 7 {
		return errors.New("message length must between 7 and 72 character")
	}

	matched, _ := regexp.MatchString(defaultTemplate, message)
	if !matched {
		return errors.New("commit message is not compatible with commit template")
	}

	log.Notice("commit message suitable for template")
	return nil
}
