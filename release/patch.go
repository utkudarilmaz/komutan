package release

import (
	"regexp"
	"sort"
	"strings"

	logging "github.com/op/go-logging"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// (?<=v\d\.\d\.)(\d)
var (
	log = logging.MustGetLogger("base")
	// patchRegexpControl = `^v\d\.\d\.\d(-(alpha|beta|teta)\.\d)?$`
	patchRegexp = `^v\d+\.\d+\.\d+$`
)

// Patch doing increase the latest tag's least significant bit to one point.
// Example: latest tag v1.0.3 -> new tag: v1.0.4
func Patch() error {

	tags, err := walkTags()
	if err != nil {
		return err
	}

	tags = filterPatchTags(tags)
	sort.Strings(tags)

	// regexp := regexp.MustCompile(patchRegexp)
	//
	// if err := regexp.MatchString(tags[len(tags)-1]); err != true {
	// 	return errors.New("no valid tag found on repo")
	// }
	//
	// regexp := regexp.MustCompile(patchRegexpControl)
	//
	log.Debug("latest tag is %s", tags[len(tags)-1])
	return nil
}

func walkTags() ([]string, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err
	}

	tagrefs, err := repo.Tags()
	if err != nil {
		return nil, err
	}

	var list []string

	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		list = append(list, strings.Trim(t.Name().String(), "refs/tags/"))
		return nil
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}

func filterPatchTags(tags []string) []string {

	filteredTags := []string{}

	regexp := regexp.MustCompile(patchRegexp)

	for i := 0; i < len(tags); i++ {
		if tmp := regexp.FindString(tags[i]); tmp != "" {
			filteredTags = append(filteredTags, tmp)
		}
	}
	return filteredTags
}
