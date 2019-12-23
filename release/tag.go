package release

import (
	"regexp"
	"strconv"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// (?<=v\d\.\d\.)(\d)
var (
	// patchRegexpControl = `^v\d\.\d\.\d(-(alpha|beta|teta)\.\d)?$`
	tagsRegexp        = `^v\d+\.\d+\.\d+$`
	tagsParsingRegexp = `^(v)(\d+)(\.)(\d+)(\.)(\d+)$`
)

type Tag struct {
	Major uint8
	Minor uint8
	Patch uint8
}

type ByVersion []Tag

func (tags ByVersion) Len() int      { return len(tags) }
func (tags ByVersion) Swap(i, j int) { tags[i], tags[j] = tags[j], tags[i] }
func (tags ByVersion) Less(i, j int) bool {
	if tags[i].Major < tags[j].Major {
		return true
	} else if tags[i].Major == tags[j].Major {
		if tags[i].Minor < tags[j].Minor {
			return true
		} else if tags[i].Patch < tags[j].Patch {
			return true
		}
	}
	return false
}

func (tag Tag) String() string {
	return ("v" +
		strconv.Itoa(int(tag.Major)) + "." +
		strconv.Itoa(int(tag.Minor)) + "." +
		strconv.Itoa(int(tag.Patch)))
}

func ParseTags(tags []string) ([]Tag, error) {
	var tagList []Tag
	regexp := regexp.MustCompile(tagsParsingRegexp)

	for i := 0; i < len(tags); i++ {
		parsedTag := regexp.FindStringSubmatch(tags[i])
		var tempTag []uint8

		for j := 2; j <= 6; j = j + 2 {

			temp, err := strconv.Atoi(parsedTag[j])
			if err != nil {
				return nil, err
			}
			tempTag[j] = uint8(temp)
		}

		tag := Tag{
			Major: tempTag[0],
			Minor: tempTag[1],
			Patch: tempTag[2],
		}
		tagList = append(tagList, tag)
	}
	return tagList, nil
}

func (tag Tag) NewRelease() error {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return err
	}

	headRef, err := repo.Head()
	if err != nil {
		return err
	}

	_, err = repo.CreateTag(tag.String(), headRef.Hash(), nil)
	if err != nil {
		return err
	}

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

func filterReleaseTags(tags []string) []string {

	filteredTags := []string{}

	regexp := regexp.MustCompile(tagsRegexp)

	for i := 0; i < len(tags); i++ {
		if tmp := regexp.FindString(tags[i]); tmp != "" {
			filteredTags = append(filteredTags, tmp)
		}
	}
	return filteredTags
}
