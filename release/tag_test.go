package release

import (
	"fmt"
	"sort"
	"testing"
)

func TestLess(t *testing.T) {
	var refList = []string{
		"v1.0.1",
		"v1.2.1",
		"v1.1.1",
		"v3.1.1",
		"v3.3.1",
		"v1.2.2",
	}

	var tagList []Tag

	tagList, err := ParseTags(refList)
	if err != nil {
		t.Errorf(err.Error())
	}

	sort.Sort(ByVersion(tagList))
	fmt.Println(tagList)

}
