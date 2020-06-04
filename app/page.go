package app

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/php-cli/models"
	"jaytaylor.com/html2text"
	"runtime"
	"sync"
)

var (
	page         models.Page
	wg           sync.WaitGroup
	examplesList []models.Example
)

func crawler(url string) {
	doc := makeARequest(url, false)
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(5)
	go readFuncDesc(doc, &page.Contents, &wg)
	go readFuncReturnValues(doc, &page.Contents, &wg)
	go readFuncExamples(doc, &examplesList, &wg)
	go readFuncNotes(doc, &page.Contents, &wg)
	go readFuncParameters(doc, &page.Contents, &wg)
	wg.Wait()
	printPageContent(page)
	PrintTable(examplesList)
}

func readFuncDesc(doc *goquery.Document, list *[]models.Content, wg *sync.WaitGroup) {
	defer wg.Done()
	content := models.Content{}
	content.Data = append(content.Data, standardizeSpaces(doc.Find(".dc-title").Text()))
	doc.Find(".description ").Each(func(i int, s *goquery.Selection) {
		content.Header = s.Find("h3").Text()
		content.Data = append(content.Data, "Usage --> "+standardizeSpaces(s.Find(".methodsynopsis").Text()))
	})

	*list = append(*list, content)
}

func readFuncReturnValues(doc *goquery.Document, list *[]models.Content, wg *sync.WaitGroup) {
	defer wg.Done()
	content := models.Content{}
	doc.Find(".returnvalues").Each(func(i int, selection *goquery.Selection) {
		content.Header = selection.Find("h3").Text()
		content.Data = append(content.Data, standardizeSpaces(selection.Find("p").Text()))
	})
	*list = append(*list, content)
}

func readFuncExamples(doc *goquery.Document, list *[]models.Example, wg *sync.WaitGroup) {
	defer wg.Done()

	doc.Find(".example").Each(func(i int, selection *goquery.Selection) {
		example := models.Example{}
		h, _ := selection.Find(".phpcode").Html()
		example.Code, _ = html2text.FromString(h, html2text.Options{PrettyTables: true})
		example.Output = selection.Find(".cdata").Text()
		*list = append(*list, example)
	})
}

func readFuncNotes(doc *goquery.Document, list *[]models.Content, wg *sync.WaitGroup) {
	defer wg.Done()
	content := models.Content{}
	doc.Find(".notes").Each(func(i int, selection *goquery.Selection) {
		content.Header = selection.Find("h3").Text()
		appendHelper(standardizeSpaces(selection.Find(".simpara").Text()), &content)
		appendHelper(standardizeSpaces(selection.Find(".para").Text()), &content)
	})
	*list = append(*list, content)
}
func readFuncParameters(doc *goquery.Document, list *[]models.Content, wg *sync.WaitGroup) {
	defer wg.Done()
	content := models.Content{}
	var data []string
	doc.Find(".parameters ").Each(func(i int, s *goquery.Selection) {

		s.Find("dt").Each(func(i int, selection *goquery.Selection) {
			data = append(data, standardizeSpaces(selection.Text()))
			//data[i] = selection.Text()
		})
		s.Find("dd").Each(func(i int, selection *goquery.Selection) {
			data[i] += "  -  " + standardizeSpaces(selection.Text())
		})

		content.Data = data
		content.Header = s.Find("h3").Text()
	})
	*list = append(*list, content)
}

func appendHelper(data string, content *models.Content) {
	if data != "" {
		content.Data = append(content.Data, data)
	}
}
