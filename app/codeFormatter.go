package app

import (
	"strings"
)

func formatCode(code string) string {
	fs := strings.NewReplacer(
		"<?php", "<?php\n",
		"?>", "?>\n",
		";", ";\n",
		"echo", "echo",
		"{", "{\n",
		//"function", "function",
		"} ", "} \n",
		"}$", "}\n$",
		"}f", "}\nf",
		"}v", "}\nv",
	).Replace(code)
	return fs
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
