package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	path := os.Args[1]
	extension := "." + os.Args[2]
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), extension) {
			os.Remove(path + "/" + file.Name())
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}
