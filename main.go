package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	root := "/Users/tamal/go/src"
	result := List(root)
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

func List(root string) []string {
	var result []string
	fs.WalkDir(os.DirFS(root), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && (d.Name() == ".go" || d.Name() == "bin" || d.Name() == "node_modules") {
			result = append(result, filepath.Join(root, path))
		}
		return nil
	})
	return result
}
