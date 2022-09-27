package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// := is shorthand for declare and assign

func extractPath(args []string) (string, error) {
	if len(args) > 0 {
		return args[0], nil
	}

	return "", errors.New("File path not given")
}

func main() {
	count := 0

	args := os.Args[1:]
	path, err := extractPath(args)

	if err != nil {
		fmt.Printf("Must provide a path to walk")
		return
	}

	err = filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			count++
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
		return
	}

	fmt.Printf("Visited: %v\n", count)
}
