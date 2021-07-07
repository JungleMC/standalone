package main

import (
	"bytes"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	println(pwd)

	matches, err := walkDirMatch(pwd, "*.proto")
	if err != nil {
		panic(err)
	}

	for _, path := range matches {
		relPath, err := filepath.Rel(pwd, path)
		if err != nil {
			panic(err)
		}

		println(relPath)
		cmd := exec.Command("protoc", "--go_out=.", "--go_opt=paths=source_relative", "--go-grpc_out=.", "--go-grpc_opt=paths=source_relative", relPath)

		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb

		err = cmd.Run()
		if err != nil {
			panic(errb.String())
		}
	}
}

func walkDirMatch(root string, pattern string) ([]string, error) {
	var matches []string
	err := filepath.WalkDir(root, func(path string, dir fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if dir.IsDir() {
			return nil
		}

		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return matches, nil
}
