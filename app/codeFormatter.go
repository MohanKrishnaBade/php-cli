package app

import (
	"strings"
)

func formatCode(code string) string {
	fs := strings.Replace(code, "<?php", "<?php\n", -1)
	fs = strings.Replace(fs, "?>", "?>\n", -1)
	fs = strings.Replace(fs, ";", ";\n", -1)
	fs = strings.Replace(fs, "echo", "\necho", -1)
	fs = strings.Replace(fs, "{", "{\n", -1)
	fs = strings.Replace(fs, "function", "\nfunction", -1)
	fs = strings.Replace(fs, "}?>", "}\n?>", -1)

	return fs
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
