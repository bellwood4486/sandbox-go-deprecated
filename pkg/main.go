package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/packages"
)

func getPackage(dir string) *packages.Package {
	p, _ := packages.Load(&packages.Config{
		Dir: dir,
	}, ".")

	if len(p) != 1 {
		return nil
	}

	return p[0]
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("WorkingDir: %s\n", wd)
	genPkg := getPackage(wd)
	fmt.Printf("Package: %s\n", genPkg.Name)
}
