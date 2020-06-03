package models

type Page struct {
	Contents []Content
}

type Content struct {
	Header string
	Data   []string
}

func NewContent() map[string][]string {
	return make(map[string][]string)
}
