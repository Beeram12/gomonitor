package checker

import (
	"fmt"
	"os"
)

func CheckStatus(urls []string) {
	fmt.Println("Enter your list of website URL's to check their status")

	// Append urls inputted into CLI into urls array

	if len(os.Args) > 1 {
		for _, i := range os.Args[1:] {
			urls = append(urls, i)
		}
	} else {
		fmt.Println("No input URL's are provided")
	}

	// For clarity print all the urls entered into the url slice
	for _, i := range urls {
		fmt.Println(i)
	}

}
