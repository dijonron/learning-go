package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed english_rights.txt
var englishRights string

//go:embed german_rights.txt
var germanRights string

//go:embed japanese_rights.txt
var japaneseRights string

func main() {
	ex1(os.Args)
}

func ex1(args []string) {
	fmt.Println("Exercise 1")

	if len(args) == 1 {
		fmt.Println("no language provided")
		return
	}

	switch args[1] {
	case "english":
		fmt.Println(englishRights)

	case "german":
		{
			fmt.Println(germanRights)
		}
	case "japanese":
		{
			fmt.Println(japaneseRights)
		}
	default:
		{
			fmt.Println("unsupported language:", args[1])

		}
	}
}
