package app

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"github.com/php-cli/models"
	"log"
	"net/http"
	"strings"
)

const (
	urlHeader     = "https://www.php.net/manual/en/"
	arrayFuncUrl  = urlHeader + "book.array.php"
	stringFuncUrl = urlHeader + "book.strings.php"
)

func Scraper(args []string) {
	list := process(args)
	if len(args) == 2 {
		funcDesc(list, args[1])
	} else {
		printList(list)
	}

}

func strategy(firstArg string) string {

	switch {
	case firstArg == "array":
		return arrayFuncUrl
	case firstArg == "string":
		return stringFuncUrl
	default:
		return urlHeader
	}
}

func process(args []string) map[string]models.Ws {
	wsList := make(map[string]models.Ws)

	// Request the HTML page.
	res, err := http.Get(strategy(args[0]))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find items
	doc.Find(".chunklist_children ").Each(func(i int, s *goquery.Selection) {

		s.Find("li").Each(func(i int, selection *goquery.Selection) {

			if strings.Contains(selection.Text(), "—") {

				ws := models.Ws{}
				selection.Find("a").Each(func(i int, selection *goquery.Selection) {
					if link, ok := selection.Attr("href"); ok {
						ws.Url = urlHeader + link
					}
				})
				d := strings.Split(selection.Text(), "—")
				ws.Name, ws.Description = d[0], d[1]

				wsList[strings.TrimSpace(ws.Name)] = ws
			}

		})
	})
	return wsList
}

func printList(list map[string]models.Ws) {
	color.Cyan.Printf("%-30s %12s\n", "Array Functions", "Description")
	println()
	for _, v := range list {
		color.FgYellow.Printf("%-30s %10s\n", v.Name, v.Description)
	}
}
