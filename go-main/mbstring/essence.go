package mbstring

import "fmt"

func Essence() {
	var name string = "Fulan"

	// Type Conversion
	var chars []byte = []byte(name)

	fmt.Printf("Output sebagai string %s\n", name)
	fmt.Printf("Output sebagai bilangan heksa % X\n", chars)
	fmt.Printf("Output sebagai bilangan desimal %d\n", chars)
	fmt.Printf("Output sebagai bilangan binner %b\n", chars)
	chars[0] = 80

	fmt.Printf("%s\n","Output Setelah Diubah :")

	fmt.Printf("Output sebagai string %s\n", string(chars))
	fmt.Printf("Output sebagai bilangan heksa % X\n", chars)
	fmt.Printf("Output sebagai bilangan desimal %d\n", chars)
	fmt.Printf("Output sebagai bilangan binner %b\n", chars)
}