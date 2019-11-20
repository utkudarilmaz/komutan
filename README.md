# Komutan

Komutan is a commit message template validator. Komutan creates a commit-msg
hook for you and be sure of your commit messages suitable the
[Conventional Commits][d1d6b24a] standartds.

## Installation

### From Source

```console
git clone github.com/utkudarilmaz/komutan
cd komutan
make install
make clean # Optional: delete the source code directory (current working directory)
```
### From tarball

1.  Download the latest binaries for release page.
2.  Untar the tar file with executing `tar -xzvf komutan.tar.gz`.
3.  Move the executable to $PATH `mv komutan /usr/local/bin/komutan`.

## Usage

Move to the directory of the repo you want to use for use Komutan and execute
`komutan init` command on the repo root directory. That is all.

### Example

--> /home/utku  
---> /home/utku/some-repo

```console
cd /home/utku/some-repo
komutan init
```

Furthermore all of your commits checked by Komutan before commit process
complete. If commit message is not valid commit process don't complete.

### Check Commit Message Validation

If you wish control your commit message before execute `git commit -m "..."`
you can execute `komutan validate -m "commit message here"` or check commit
message from file `komutan validate -f <path_of_commit_message>`.

### Delete the Hook Created by Komutan

If you want remove the commit message hook from your repository you just execute
`rm -f <your_repo's_root_path>/.git/hooks/commit-msg`. That is all.

#### Example:

```console
cd some-repository
rm -f .git/hooks/commit-msg
```

## Commit Schema

```
TYPE[SUB-TYPE]: DESCRIPTION

[BODY]

[FOOTER]
```

You can find more details on [Conventional Commits][d1d6b24a]
about commit message rules.

### Rules

1. Commit message's header section (TYPE[SUB-TYPE]: DESCRIPTION) can't be more
   than 72 character.
2. Description's first character must be one of [a-z-.].
3. Description's last character can't be one of . , ! SPACE ? {} []

### Available Types

+   feat
+   docs
+   style
+   perf
+   test
+   fix
+   refactor
+   chore

[d1d6b24a]: https://www.conventionalcommits.org "Conventional Commits"
