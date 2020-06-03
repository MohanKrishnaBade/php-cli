package app

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"github.com/php-cli/models"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

var page models.Page
var wg sync.WaitGroup

func crawler(url string) {
	// Request the HTML page.
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

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(5)
	go readDesc(doc, &page.Contents, &wg)
	go readReturnValues(doc, &page.Contents, &wg)
	go readExamples(doc, &page.Contents, &wg)
	go readNotes(doc, &page.Contents, &wg)
	go readParameters(doc, &page.Contents, &wg)
	wg.Wait()
	printer(page)
}

func readDesc(doc *goquery.Document, list *[]models.Content, wg *sync.WaitGroup) {
	defer wg.Done()
	content := models.Content{}
	doc.Find(".description ").Each(func(i int, s *goquery.Selection) {
		content.Header = s.Find("h3").Text()
		content.Data = append(content.Data, standardizeSpaces(s.Find(".methodsynopsis").Text()))
	})
	*list = append(*list, content)
}

func readReturnValues(doc *goquery.Document, list *[]models.Content, wg *sync.WaitGroup) {
	defer wg.Done()
	content := models.Content{}
	doc.Find(".returnvalues").Each(func(i int, selection *goquery.Selection) {
		content.Header = selection.Find("h3").Text()
		content.Data = append(content.Data, standardizeSpaces(selection.Find("p").Text()))
	})
	*list = append(*list, content)
}

func readExamples(doc *goquery.Document, list *[]models.Content, wg *sync.WaitGroup) {
	defer wg.Done()
	content := models.Content{}
	doc.Find(".examples").Each(func(i int, selection *goquery.Selection) {
		content.Header = selection.Find("h3").Text()
		content.Data = append(content.Data, standardizeSpaces(selection.Find("span").Text()))
	})
	*list = append(*list, content)
}

func readNotes(doc *goquery.Document, list *[]models.Content, wg *sync.WaitGroup) {
	defer wg.Done()
	content := models.Content{}
	doc.Find(".notes").Each(func(i int, selection *goquery.Selection) {
		content.Header = selection.Find("h3").Text()
		content.Data = append(content.Data, selection.Find(".para").Text())
	})
	*list = append(*list, content)
}
func readParameters(doc *goquery.Document, list *[]models.Content, wg *sync.WaitGroup) {
	defer wg.Done()
	content := models.Content{}
	var data []string
	doc.Find(".parameters ").Each(func(i int, s *goquery.Selection) {

		s.Find("dt").Each(func(i int, selection *goquery.Selection) {
			data = append(data, standardizeSpaces(selection.Text()))
			//data[i] = selection.Text()
		})
		s.Find("dd").Each(func(i int, selection *goquery.Selection) {
			data[i] += "  --  " + standardizeSpaces(selection.Text())
			//content.Data = append(content.Data, "____"+standardizeSpaces(selection.Text()))

		})

		content.Data = data
		content.Header = s.Find("h3").Text()
	})
	*list = append(*list, content)
}
func funcDesc(list map[string]models.Ws, key string) {
	if v, ok := list[key]; ok {
		color.FgYellow.Printf("%s---->%s\n", v.Name, v.Description)
		crawler(v.Url)
	}
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func printer(page models.Page) {

	for _, v := range page.Contents {
		color.FgGreen.Printf("%-30s\n", v.Header)
		color.FgLightWhite.Println(strings.Repeat("-", len(v.Header)+1))
		for _, line := range v.Data {
			if v.Header == "Examples" {
				Format(line)
			} else {
				fmt.Printf("%s%s\n", "●  ", line)
			}
		}
		fmt.Println()
	}
}

func prettier(str string, format string, length int) {
	for len(str) > length {
		fmt.Printf(format, str[:length])
		str = str[length:]
	}
	fmt.Printf(format, str)
}
func data() {
	//content := models.NewContent()
	//var header []string

	//function description part
	//doc.Find(".description ").Each(func(i int, s *goquery.Selection) {
	//	header = append(header, s.Find("h3").Text())
	//	data := append([]string{}, standardizeSpaces(s.Find(".methodsynopsis").Text()))
	//	content[header[0]] = data
	//
	//})

	//var data []string
	//doc.Find(".parameters ").Each(func(i int, s *goquery.Selection) {
	//
	//	s.Find("dt").Each(func(i int, selection *goquery.Selection) {
	//		data = append(data, selection.Text())
	//		//data[i] = selection.Text()
	//	})
	//	s.Find("dd").Each(func(i int, selection *goquery.Selection) {
	//		data[i] += "--" + standardizeSpaces(selection.Text())
	//	})
	//
	//	//data := append([]string{}, standardizeSpaces(s.Find("dl").Text()))
	//	header = append(header, s.Find("h3").Text())
	//	content[header[1]] = data
	//})

	////function return values
	//doc.Find(".returnvalues").Each(func(i int, selection *goquery.Selection) {
	//	header = append(header, selection.Find("h3").Text())
	//	data := append([]string{}, standardizeSpaces(selection.Find("p").Text()))
	//	content[header[2]] = data
	//})
	//
	////function usage examples
	//doc.Find(".examples").Each(func(i int, selection *goquery.Selection) {
	//	header = append(header, selection.Find("h3").Text())
	//	data := append([]string{}, selection.Find("span").Text())
	//	content[header[3]] = data
	//})
	//
	////function notes
	//doc.Find(".notes").Each(func(i int, selection *goquery.Selection) {
	//	header = append(header, selection.Find("h3").Text())
	//	data := append([]string{}, selection.Find(".para").Text())
	//	content[header[4]] = data
	//})
	//
	//page.Headers = header
	//page.Content = content

	//for k, v := range page.Content {
	//	fmt.Println(k, v)
	//}
	//color.FgCyan.Println(page.Content)

}
