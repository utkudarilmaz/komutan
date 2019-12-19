package release

import (
	"regexp"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// (?<=v\d\.\d\.)(\d)
var (
	// patchRegexpControl = `^v\d\.\d\.\d(-(alpha|beta|teta)\.\d)?$`
	tagsRegexp = `^v\d+\.\d+\.\d+$`
)

// Patch doing increase the latest tag's least significant bit to one point.
// Example: latest tag v1.0.3 -> new tag: v1.0.4

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

	regexp := regexp.MustCompile(tagsRegexp)

	for i := 0; i < len(tags); i++ {
		if tmp := regexp.FindString(tags[i]); tmp != "" {
			filteredTags = append(filteredTags, tmp)
		}
	}
	return filteredTags
}
