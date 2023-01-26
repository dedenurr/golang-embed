package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"

	"testing"
)

//1. Embed text 1 file
//go:embed version1.txt

var b string
func TestString(t *testing.T) {
	fmt.Println(b)
}


//2. embed file gambar
//go:embed logo.png
var logo []byte //untuk embed file gambar

func TestByteArray(t *testing.T)  {
	err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	
}


// 3. embed multiple file
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T)  {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}


// 4. patch matcher untuk membaca multiple file yang kita inginkan
//go:embed files/*.txt
var path embed.FS

func TestPatchMatcher(t *testing.T)  {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries{
		if !entry.IsDir(){
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content: ",string(content))
		}
	}
}