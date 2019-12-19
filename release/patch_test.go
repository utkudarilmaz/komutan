package release

import (
	"os"
	"testing"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// Patch release test function
func TestPatch(t *testing.T) {

	tmpDirName := "/tmp/komutan"
	tmpFileName := "/tmp/komutan/file"

	tags := []string{"v0.0.1", "v1.0.1", "v1.1.1", "v2.0.3", "v3.0.1-alpha.0"}

	err := os.Mkdir(tmpDirName, 0700)
	if err != nil {
		t.Errorf(
			"test repo directory can't be created at %s",
			tmpDirName,
		)
	}
	_, err = os.Create(tmpFileName)
	if err != nil {
		t.Errorf(
			"test file can't be created at %s",
			tmpFileName,
		)
	}

	repo, err := git.PlainInit(tmpDirName, false)
	if err != nil {
		t.Errorf(err.Error())
	}

	workTree, err := repo.Worktree()
	if err != nil {
		t.Errorf(err.Error())
	}
	_, _ = workTree.Add(tmpDirName)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// }

	commitOptions := &git.CommitOptions{
		Author: &object.Signature{Name: "Foo", Email: "foo@example.local", When: time.Now()},
	}
	hash, err := workTree.Commit("test", commitOptions)
	if err != nil {
		t.Errorf(err.Error())
	}

	for i := 0; i < len(tags); i++ {
		_, err := repo.CreateTag(tags[i], hash, nil)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	err = os.Chdir(tmpDirName)
	if err != nil {
		t.Errorf(err.Error())
	}

	bit, err := Patch()
	if err != nil {
		t.Errorf(err.Error())
	}
	if bit != "3" {
		t.Errorf(
			"expected patch tag is 3 but found %s",
			bit,
		)
	}
	err = os.RemoveAll(tmpDirName)
	if err != nil {
		t.Errorf(err.Error())
	}

}
