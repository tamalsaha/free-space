package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	dir := "/Users/tamal/go/src"
	result := List(os.DirFS(dir))
	sort.Strings(result)

	remaining := make([]string, 0, len(result))
	for _, entry := range result {
		err := os.RemoveAll(entry)
		if err != nil {
			remaining = append(remaining, entry)
		} else {
			fmt.Println("rmoved:", entry)
		}
	}

	fmt.Println("----------------------------------------------")
	for _, entry := range remaining {
		fmt.Println("sudo rm -rf", entry)
	}
}

func List(fsys fs.FS) []string {
	var result []string
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && (d.Name() == ".go" || d.Name() == "bin" || d.Name() == "node_modules") {
			result = append(result, filepath.Join(path, d.Name()))
		}
		return nil
	})
	return result
}
