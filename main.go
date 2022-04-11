package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/bep/golibsass/libsass"
)

func main() {

	transpiler, _ := libsass.New(libsass.Options{
		OutputStyle:      libsass.ExpandedStyle,
		IncludePaths:     []string{"htdocs/themes/site/src/sass/"},
		SourceMapOptions: libsass.SourceMapOptions{EnableEmbedded: true},
	})

	content, err := ioutil.ReadFile("htdocs/themes/site/src/sass/email.scss")
	// content, err := ioutil.ReadFile("htdocs/themes/site/src/sass/test.css")
	if err != nil {
		log.Fatal(err)
	}

	result, err := transpiler.Execute(string(content))
	if err != nil {
		panic(err)
	}

	fmt.Println(result.CSS)
	fmt.Println(result.SourceMapContent)
}
