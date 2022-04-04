package main

import (
	"fmt"

	"github.com/louissaadgo/go-checkif"
)

func main() {

	phrase := "checkif is so powerful!"

	phraseChecker := checkif.StringObject{
		Data: phrase,
	}
	phraseChecker.IsShorterThan(20).IsLongerThan(3)

	fmt.Println(phraseChecker)
}
