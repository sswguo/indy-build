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

func GoGit(gitURL, dir, checkout, checkoutType string) {
	// validate and refine checkout target, tag vs branch, and default to branch:master
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

	directory := fmt.Sprintf("/tmp/%s", dir)
	clone := true
	if fileExists(directory) {
		clone = false
	}

	// Cloning or opening exists repo
	var r *git.Repository
	var err error
	if clone {
		fmt.Printf("Start cloning %s to %s\n", gitURL, directory)
		r, err = git.PlainClone(directory, false, &git.CloneOptions{
			URL:               gitURL,
			Progress:          os.Stdout,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Open existed repo %s\n", directory)
		r, err = git.PlainOpen(directory)
	}

	w, err := r.Worktree()

	if err != nil {
		log.Fatal(err)
	}

	// Updating heads
	fmt.Printf("Fetching Updates....\n")
	r.Fetch(&git.FetchOptions{
		RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
		Progress: os.Stdout,
	})

	// Checking out to branch or tag
	fmt.Printf("Checking out to %s\n", checkoutFull)
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(checkoutFull),
		Force:  true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Show HEAD
	ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Now HEAD is %s\n", ref.Hash())

	// defer os.RemoveAll(directory)
}
