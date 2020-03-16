package main

import (
	"fmt"
	"theory/dict/dicts"
)

func main() {
  dictionary := dicts.Dictionary{"first": "World"}
  dictionary.Search("second")
  fmt.Println()
}
