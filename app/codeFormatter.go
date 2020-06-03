package app

import (
	"github.com/gookit/color"
	"strings"
)

// strrev() <?phpecho strrev("Hello world!"); // outputs "!dlrow olleH"?> <?phpecho strrev("Hello world!"); // outputs "!dlrow olleH"?>

func Format(code string) {
	d := strings.Replace(code, "<?php", "\n<?php\n", -1)
	d = strings.Replace(d, "?>", "?>\n", -1)
	d = strings.Replace(d, ";", ";\n", -1)
	d = strings.Replace(d, "echo", "\necho", -1)
	d = strings.Replace(d, "{", "{\n", -1)
	d = strings.Replace(d, "function", "\nfunction", -1)
	color.FgDarkGray.Println(d)
}
