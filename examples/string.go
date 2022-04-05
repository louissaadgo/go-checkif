package main

import (
	"fmt"

	"github.com/louissaadgo/go-checkif"
)

func main() {

	phrase := "kfDk"

	phraseChecker := checkif.StringObject{
		Data: phrase,
	}
	phraseChecker.DoesNotContainUpperCaseLetter()

	fmt.Println(phraseChecker)
}
