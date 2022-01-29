package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)

	if err != nil {
		panic(err.Error())
	}

	//if err != nil {
	//	panic(err.Error())
	//}

	//fmt.Println(os.Args)

	//argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]
	//
	//arg := os.Args[3]
	//fmt.Println(argsWithProg)
	//fmt.Println(argsWithoutProg)
	//fmt.Println(arg)

	//fmt.Println(path)
	//fmt.Println(printFiles)
	//
	//files, err := ioutil.ReadDir(path)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, file := range files {
	//	fmt.Println(file.Name(), file.IsDir())
	//}
}

func dirTree(out *os.File, path string, printFiles bool) error {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {

		if !file.IsDir() && printFiles {
			out.WriteString(file.Name() + strconv.Itoa(file.Size()))
			out.WriteString("\n")
		} else if file.IsDir() {
			out.WriteString(file.Name())
			out.WriteString("\n")
		}
	}

	fmt.Fprintln(out, "\n")

	return nil
}
