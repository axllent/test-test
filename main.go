package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("Usage: %s <file.scss>\n", args[0])
		os.Exit(1)
	}

	file := args[1]

	fmt.Println(file)

	// base, err := filepath.Abs(path.Dir(file))
	// if err != nil {
	// 	panic(err)
	// }

	// transpiler, _ := libsass.New(libsass.Options{
	// 	OutputStyle:      libsass.ExpandedStyle,
	// 	IncludePaths:     []string{base},
	// 	SourceMapOptions: libsass.SourceMapOptions{EnableEmbedded: true},
	// })

	// content, err := ioutil.ReadFile(file)
	// if err != nil {
	// 	panic(err)
	// }

	// result, err := transpiler.Execute(string(content))
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(result.CSS)
	// fmt.Println(result.SourceMapContent)
}
