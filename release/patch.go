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
	patchRegexp = `^v\d\.\d\.\d$`
)

type alphabetic []string

func (list alphabetic) Len() int { return len(list) }

func (list alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list alphabetic) Less(i, j int) bool {
	var si string = list[i]
	var sj string = list[j]
	return si < sj
}

func Patch() error {

	tags, err := walkTags()
	if err != nil {
		return err
	}

	filterPatchTags(tags)
	sort.Sort(alphabetic(tags))

	// regexp := regexp.MustCompile(patchRegexp)
	//
	// if err := regexp.MatchString(tags[len(tags)-1]); err != true {
	// 	return errors.New("no valid tag found on repo")
	// }
	//
	// regexp := regexp.MustCompile(patchRegexpControl)
	//
	log.Debug("latest tag is %s", tags)
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

func filterPatchTags(tags []string) {

	filteredTags := []string{}

	regexp := regexp.MustCompile(patchRegexp)
	for i := 0; i < len(tags); i++ {
		filteredTags = append(filteredTags, regexp.FindString(tags[i]))
	}
	copy(tags, filteredTags)

}
