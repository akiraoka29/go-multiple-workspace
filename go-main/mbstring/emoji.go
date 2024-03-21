package mbstring

import "fmt"

func Emoji() {
	// var emoji rune = '😝'
	// var emoji string = "😝"
	var symbol string = "😝😝😝"
	var emoji []rune = []rune(symbol)

	// With Rune
	// fmt.Printf("Output sebagai string %s\n", string(emoji))
	// fmt.Printf("Output sebagai rune atau unicode %U\n", emoji)
	// fmt.Printf("Output sebagai bilangan heksa % X\n", emoji)
	// fmt.Printf("Output sebagai bilangan desimal %d\n", emoji)
	// fmt.Printf("Output sebagai bilangan binner %b\n", emoji)

	// With String
	fmt.Printf("Output sebagai string %s\n", string(emoji))
	fmt.Printf("Output sebagai rune atau unicode %U\n", []rune(emoji))
	fmt.Printf("Output sebagai bilangan heksa % X\n", []rune(emoji))
	fmt.Printf("Output sebagai bilangan desimal %d\n", []rune(emoji))
	fmt.Printf("Output sebagai bilangan binner %b\n", []rune(emoji))
}

func EmojiUnicode() {
	var emoji string = "\U0001F61D \t \U0001F61D \t \U0001F61D" // Unicode with 8 Character atau Escape Sequence
	// var emoji string = "\u2639 \u2639 \u2639" // Unicode with 8 Character atau Escape Sequence

	fmt.Printf("Output sebagai string %s\n", string(emoji))
	fmt.Printf("Output sebagai rune atau unicode %U\n", []rune(emoji))
	fmt.Printf("Output sebagai bilangan heksa % X\n", []rune(emoji))
	fmt.Printf("Output sebagai bilangan desimal %d\n", []rune(emoji))
	fmt.Printf("Output sebagai bilangan binner %b\n", []rune(emoji))
}