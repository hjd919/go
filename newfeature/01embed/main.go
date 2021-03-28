package main

import (
	"embed"
	"fmt"
)

//go:embed testdir
var f embed.FS

func main() {
	dirpath := "testdir"
	dirs, err := f.ReadDir(dirpath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, dir := range dirs {
		cBytes, err := f.ReadFile(dirpath + "/" + dir.Name())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(cBytes))
	}
}
