package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	h             bool
	initDev       bool
	clear         bool
	dirName       string
	mainGoContent = `package main

import ()

func main() {

}
	`
)

// parse parameter

func ParseFlag() {
	flag.BoolVar(&h, "h", false, "show help")
	flag.BoolVar(&initDev, "init", false, "init a development for trying golang std module or trying open source code's quickstart \n just a new folder contains gomod file in local")
	flag.BoolVar(&clear, "clear", false, "clear development")
	flag.StringVar(&dirName, "n", "trygo", "name for dirname")
	flag.Parse()

}
func isPathNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func main() {
	ParseFlag()
	var goModContent = fmt.Sprintf("module %s \n\ngo 1.13", dirName)
	switch {
	case h:
		fmt.Println("init: init a go mod dev in local \nclear: clear local path")

	case initDev:

		// path join
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dirPath := filepath.Join(pwd, dirName)
		fmt.Println("dirPath: ", dirPath)
		// generate folder
		notExist, err := isPathNotExist(dirPath)
		if err != nil {
			fmt.Println("judge isPathExist error, error dirPath:", dirPath, " err info:", err)
			os.Exit(1)
		}
		// generate go.mod & main.go
		if notExist {
			err := os.Mkdir(dirPath, os.ModePerm)
			if err != nil {
				if os.IsPermission(err) {
					fmt.Println("mkdir dir path: ", dirPath, " no permission")
					os.Exit(1)

				}
				fmt.Println(err)
				os.Exit(1)
			}
		}
		//fmt.Println("dirPath: ", dirPath)
		//err = os.Chdir(dirPath)
		mainf, err := os.Create(dirPath + "/main.go")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer mainf.Close()
		mod, err := os.Create(dirPath + "/go.mod")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer mod.Close()
		_, err = io.WriteString(mod, goModContent)
		_, err = io.WriteString(mainf, mainGoContent)

	case clear:
		fmt.Println("clear")
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dirPath := filepath.Join(pwd, dirName)
		err = os.RemoveAll(dirPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println()
	}
	// TODO auto cd dir

}
