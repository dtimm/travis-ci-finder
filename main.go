package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {
	repo := os.Args[1]
	fs := memfs.New()
	_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL: repo,
	})
	if err != nil {
		panic(err)
	}

	_, err = fs.Open("/.travis.yml")
	if err != nil {
		fmt.Printf("No .travis.yml file found in %s: %s\n", repo, err)
		os.Exit(0)
	}
	fmt.Printf("Found .travis.yml file in %s\n", repo)
}
