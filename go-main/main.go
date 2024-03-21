package main

import (
	"fmt"

	"github.com/akiraoka29/belajar-multiple-workspace/mbstring"
)

func main() {
	// Type Inference
	format := "%s %s\n"
	fmt.Printf(format,"Hello", "World")

	var name string = "Fulan"
	var setName *string = &name // & = Pass by Reference / Pass By Memory. * Tipe Data Yang Mempointer Langsung ke Memory
	fmt.Printf("%p\n", setName)
	fmt.Printf("%s\n", *setName) // * Diawal Variable, Itu Namanya Deference Pointer

	// main2()

	// mbstring.Essence()
	// mbstring.Emoji()
	mbstring.EmojiUnicode()
}