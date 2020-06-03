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

//  method that initiates the process
func Scraper(args []string) {
	list := parseHtmlData(makeARequest(args[0], true))
	if len(args) == 2 {

		if v, ok := list[args[1]]; ok {
			crawler(v.Url)
		} else {
			color.FgRed.Printf("%s(%s)\n", "we couldn't find the function name.", args[1])
		}
	} else {
		printList(list)
	}
}

// strategy methods to decide which one to call
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

// prepare a list from html data
func parseHtmlData(data *goquery.Document) map[string]models.Ws {
	wsList := make(map[string]models.Ws)
	data.Find(".chunklist_children ").Each(func(i int, s *goquery.Selection) {

		s.Find("li").Each(func(i int, selection *goquery.Selection) {

			if strings.Contains(selection.Text(), "—") {

				ws := models.Ws{}
				selection.Find("a").Each(func(i int, selection *goquery.Selection) {
					if link, ok := selection.Attr("href"); ok {
						ws.Url = urlHeader + link
					}
				})
				d := strings.Split(selection.Text(), "—")
				ws.Name, ws.Description = d[0], standardizeSpaces(d[1])

				wsList[strings.TrimSpace(ws.Name)] = ws
			}

		})
	})

	return wsList
}

// this function is responsible for making http requests.
func makeARequest(url string, useStrategy bool) *goquery.Document {
	if useStrategy {
		url = strategy(url)
	}
	res, err := http.Get(url)
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
	return doc
}
