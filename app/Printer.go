package app

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/php-cli/models"
	"os"
	"strings"
)

// this method is responsible for printing the entire list of array/string functions
func printList(list map[string]models.Ws) {
	println()
	for _, v := range list {
		color.FgYellow.Printf("%-30s %10s\n", v.Name, v.Description)
	}
	color.FgGreen.Printf("\n(%v)  %s \n", len(list), " build in functions are available in PHP")
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
				fmt.Printf("%s%s\n", "●  ", line)
			}
			fmt.Println()
		}
	}
}

func PrintTable(data []models.Example) {
	if len(data) != 0 {
		color.FgGreen.Printf("%-1s(%v)\n", "Examples", len(data))
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Example Code", "output"})
		for k, v := range data {
			if v.Output == "" {
				v.Output = fmt.Sprintf("%v  %v", emoji.SlightlySmilingFace, emoji.ManTechnologist)
			}
			if k == 0 {
				t.AppendRows([]table.Row{
					{v.Code, v.Output},
				})
			} else {
				t.AppendRow([]interface{}{v.Code, v.Output})
			}
		}

		t.SetStyle(table.StyleLight)
		t.Style().Format.Footer = text.FormatLower
		t.Style().Options.DrawBorder = true
		t.Style().Options.SeparateRows = true
		t.Render()

	}

	//t.SetStyle(table.StyleLight)
	//t.Style().Options.SeparateRows=true
	//t.Style().Options.DrawBorder = false
}
