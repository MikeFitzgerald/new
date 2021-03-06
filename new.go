package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var attributes = []byte(`# All text files should have the "lf" (Unix) line endings
* text eol=lf

# Explicitly declare text files you want to always be normalized and converted
# to native line endings on checkout.
*.java text
*.js text
*.css text
*.html text

# Denote all files that are truly binary and should not be modified.
*.png binary
*.jpg binary
*.pdf binary
*.eot binary
*.ttf binary
*.gzip binary
*.gz binary
*.ai binary
*.eps binary
*.swf binary
*.gif binary
*.mp4 binary
*.dll binary
*.class binary
*.exe binary
*.ico binary
*.doc binary`)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	args := os.Args[1:]

	switch args[0] {
	case "project":
		project(args[1])
	default:
		fmt.Println("I'm not sure what you want to do.")
	}

}

// Create a new programming project
func project(name string) {
	data := []byte("# " + properTitle(name))
	//empty := []byte("")
	//files := []string{".gitattributes", ".gitignore"}

	// Create a new directory with the project's name
	// Create a README.md
	// Create a .gitattributes file
	// Create a .gitignore file
	err := os.Mkdir(name, 0644)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(name+"/README.md", data, 0644)
	check(err)
	err = ioutil.WriteFile(name+"/.gitattributes", attributes, 0644)
	check(err)
	err = ioutil.WriteFile(name+"/.gitignore", []byte(""), 0644)
	check(err)
}

func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := "a an on the to"
	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
