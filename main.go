package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	dir   string
	depth int
	h     bool
)

func init() {
	flag.StringVar(&dir, "path", ".", "basic path")
	flag.IntVar(&depth, "depth", -1, "max depth, -1 is unrestricted")
	flag.BoolVar(&h, "h", false, "help")

	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		subPath := path[len(dir):]
		if len(subPath) == 0 {
			printPath(path, 0)
			return nil
		}
		arr := strings.Split(subPath, string(filepath.Separator))
		skip := 0
		if subPath[0] == filepath.Separator {
			skip = 1
		}
		if depth >= 0 && len(arr)-skip > depth {
			return nil
		}
		printPath(path, len(arr)-skip)
		return nil
	})
}

func printPath(p string, level int) {
	base := filepath.Base(p)
	if level == 0 {
		fmt.Println(base)
	} else {
		for i := 0; i < level-1; i++ {
			fmt.Printf("|  ")
		}
		fmt.Println("|-", base)
	}
}
