package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printTree(path string, prefix string, isLast bool, depth int) {
	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	if isLast {
		fmt.Println(path)
	}
	for i, entry := range entries {
		if entry.IsDir() && depth > 0 {
			if i == len(entries)-1 {
				fmt.Println(prefix, "└──", entry.Name())
				printTree(path+"\\"+entry.Name(), prefix+"	", false, depth-1)
			} else {
				fmt.Println(prefix, "│──", entry.Name())
				printTree(path+"\\"+entry.Name(), prefix+" │	", false, depth-1)
			}
		} else {
			if i == len(entries)-1 {
				fmt.Println(prefix, "└──", entry.Name())
			} else {
				fmt.Println(prefix, "│──", entry.Name())
			}
		}
	}
}

func main() {
	var depth int
	flag.IntVar(&depth, "n", 2, "Глубина дерева")
	flag.Parse()
	path := flag.Arg(0)
	if !strings.HasPrefix(path, "\\") {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = filepath.Join(wd, path)
	}
	printTree(path, "", true, depth)
}
