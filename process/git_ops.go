package process

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func getGitDir(gitURL string) string {
	segs := strings.Split(gitURL, "/")
	last := segs[len(segs)-1]
	if strings.Contains(last, ".git") {
		splits := strings.Split(last, ".")
		return splits[0]
	}
	return last
}

func GetSrc(gitURL, checkout, checkoutType string) string {
	dir := fmt.Sprintf("%s/%s", getTempDir(), getGitDir(gitURL))
	return goGit(gitURL, dir, checkout, checkoutType)
}

func goGit(gitURL, directory, checkout, checkoutType string) string {
	clone := true
	if fileExists(directory) {
		clone = false
	}

	r, err := getRepo(gitURL, directory, clone)
	checkIfError(err)

	// Updating heads
	fetchUpdates(r)

	// Checking out to branch or tag
	checkoutTo(r, getCheckout(checkout, checkoutType))

	showHEAD(r)

	return directory
}

// validate and refine checkout target, tag vs branch, and default to branch:master
func getCheckout(checkout, checkoutType string) string {
	cType := strings.TrimSpace(checkoutType)
	if strings.TrimSpace(checkout) == "" {
		checkout = "master"
	}
	var checkoutFull string
	if cType == "tag" {
		checkoutFull = fmt.Sprintf("refs/tags/%s", checkout)
	} else {
		checkoutFull = fmt.Sprintf("refs/heads/%s", checkout)
	}
	return checkoutFull
}

// Cloning or opening exists repo
func getRepo(gitURL, directory string, clone bool) (r *git.Repository, err error) {
	if clone {
		fmt.Printf("Start cloning %s to %s\n", gitURL, directory)
		r, err = git.PlainClone(directory, false, &git.CloneOptions{
			URL:               gitURL,
			Progress:          os.Stdout,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
	} else {
		fmt.Printf("Open existed repo %s\n", directory)
		r, err = git.PlainOpen(directory)
	}
	return r, err
}

// Fetching updates...
func fetchUpdates(r *git.Repository) {
	fmt.Printf("Fetching Refs....\n")
	err := r.Fetch(&git.FetchOptions{
		RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
		Progress: os.Stdout,
	})
	if err != git.NoErrAlreadyUpToDate {
		checkIfError(err)
	}
}

// Checking out to branch or tag
func checkoutTo(r *git.Repository, checkout string) {
	w, err := r.Worktree()
	checkIfError(err)

	fmt.Printf("Checking out to %s\n", checkout)
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(checkout),
		Force:  true,
	})
	checkIfError(err)
}

func showHEAD(r *git.Repository) {
	// Show HEAD
	ref, err := r.Head()
	checkIfError(err)
	fmt.Printf("Now HEAD is %s\n", ref.Hash())

	commitIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	checkIfError(err)

	commit, err := commitIter.Next()
	checkIfError(err)
	hash := commit.Hash.String()
	line := strings.Split(commit.Message, "\n")
	fmt.Println(hash[:7], line[0])
}

func checkIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func rmRepo(repoDir string) {
	fmt.Printf("Cleaning repo dir %s\n", repoDir)
	os.RemoveAll(repoDir)
}
