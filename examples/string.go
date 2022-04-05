package main

import (
	"fmt"

	"github.com/louissaadgo/go-checkif"
)

func main() {

	phrase := "swwaecd"

	phraseChecker := checkif.StringObject{
		Data: phrase,
	}
	phraseChecker.DoesNotContainCustomString("wwa")

	fmt.Println(phraseChecker)
}
