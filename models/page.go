package models

type Page struct {
	Contents []Content
	Examples []Example
}

type Content struct {
	Header string
	Data   []string
}

type Example struct {
	Code   string
	Output string
}
