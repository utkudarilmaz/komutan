package release

import (
	"errors"
	"regexp"
	"sort"

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
func Patch() (string, error) {

	tags, err := walkTags()
	if err != nil {
		return "", err
	}

	tags = filterPatchTags(tags)
	if len(tags) < 1 {
		return "", errors.New("Any tag found on repository")
	}
	sort.Strings(tags)
	bit, err := findLatestPatchBit(tags[len(tags)-1])
	if err != nil {
		return "", err
	}

	log.Debug("latest tag is %s", tags[len(tags)-1])
	return bit, nil
}

func InitPatch() {
	// TODO: create initial patch tag
}

func findLatestPatchBit(tag string) (string, error) {
	regexp := regexp.MustCompile(patchRegexp)
	var bit []string
	if bit = regexp.FindStringSubmatch(tag); len(bit) < 1 {
		return "", errors.New("No available tags found")
	}

	return bit[2], nil
}
