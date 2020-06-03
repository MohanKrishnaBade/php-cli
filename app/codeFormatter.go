package app

import (
	"github.com/gookit/color"
	"strings"
)

func formatCode(code string) {
	d := strings.Replace(code, "<?php", "\n<?php\n", -1)
	d = strings.Replace(d, "?>", "?>\n", -1)
	d = strings.Replace(d, ";", ";\n", -1)
	d = strings.Replace(d, "echo", "\necho", -1)
	d = strings.Replace(d, "{", "{\n", -1)
	d = strings.Replace(d, "function", "\nfunction", -1)
	d = strings.Replace(d, "}?>", "}\n?>", -1)
	color.FgDarkGray.Println(d)
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
