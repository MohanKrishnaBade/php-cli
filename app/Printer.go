package app

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/php-cli/models"
	"strings"
)

// this method is responsible for printing the entire list of array/string functions
func printList(list map[string]models.Ws) {
	color.Cyan.Printf("%-30s %12s\n", "Array Functions", "Description")
	println()
	for _, v := range list {
		color.FgYellow.Printf("%-30s %10s\n", v.Name, v.Description)
	}
}

func prettier(str string, format string, length int) {
	for len(str) > length {
		fmt.Printf(format, str[:length])
		str = str[length:]
	}
	fmt.Printf(format, str)
}

// this will print page content in a well formatted way
func printPageContent(page models.Page) {

	for _, v := range page.Contents {
		if len(v.Data) > 0 {
			color.FgGreen.Printf("%-30s\n", v.Header)
			color.FgLightWhite.Println(strings.Repeat("-", len(v.Header)+1))
			for _, line := range v.Data {
				if v.Header == "Examples" {
					formatCode(line)
				} else {
					fmt.Printf("%s%s\n", "●  ", line)
				}
			}
			fmt.Println()
		}
	}
}
