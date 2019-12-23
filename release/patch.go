package release

import (
	"errors"
	"regexp"
	"sort"
	"strconv"

	logging "github.com/op/go-logging"
)

// (?<=v\d\.\d\.)(\d)
var (
	log = logging.MustGetLogger("base")
	// patchRegexpControl = `^v\d\.\d\.\d(-(alpha|beta|teta)\.\d)?$`
	patchRegexp = `^(v\d+\.\d+\.)(\d+)$`
)

// Patch doing increase the latest tag's least significant bit to one point.
// Example: latest tag v1.0.3 -> new tag: v1.0.4
func Patch() error {

	tags, err := walkTags()
	if err != nil {
		return err
	}

	tags = filterReleaseTags(tags)
	if len(tags) < 1 {
		return errors.New("Any tag found on repository")
	}
	sort.Strings(tags)
	log.Debug("latest tag is %s", tags[len(tags)-1])

	parsedLatestTag, err := findLatestPatchBit(tags[len(tags)-1])
	if err != nil {
		return err
	}

	patchedBit, err := strconv.Atoi(parsedLatestTag[2])

	if err != nil {
		return err
	}

	patchedBit += 1
	tagName := parsedLatestTag[1] + strconv.Itoa(patchedBit)

	// err = newTag(tagName)
	// if err != nil {
	// 	return err
	// }
	log.Notice("%s named tag created", tagName)
	log.Notice("git push origin %s", tagName)

	return nil
}

func InitPatch() {
	// TODO: create initial patch tag
}

func findLatestPatchBit(tag string) ([]string, error) {
	regexp := regexp.MustCompile(patchRegexp)
	var parsedLatestTag []string
	if parsedLatestTag = regexp.FindStringSubmatch(tag); len(parsedLatestTag) < 1 {
		return nil, errors.New("No available tags found")
	}

	return parsedLatestTag, nil
}
