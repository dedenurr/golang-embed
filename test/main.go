package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version1.txt
var version1 string

//go:embed logo.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version1)

	err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries{
		if !entry.IsDir(){
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content: ",string(content))
		}
	}
}