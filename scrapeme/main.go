package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

var listOfUrls []string
var listOfProcessedUrls []string
var url string
var sortedResults map[int][]string

func saveBasedOnStatusCode(myUrl string, code int) {
	if _, err := myUrl[code]; err != nil {
		myUrl[code] = make([]string, 0)
	}
	sortedResults[code] = append(sortedResults[code], myUrl)
}

func collectURLOutput() {
	for _, myUrl := range listOfUrls {
		if slices.Contains(listOfProcessedUrls, myUrl) {
			continue
		}

		if !strings.Contains(myUrl, url) {
			continue
		}

		resp, err := http.Get(myUrl)
		fmt.Printf("### [PROCESSING][%d]: %s\n", resp.StatusCode, myUrl)
		if err != nil {
			fmt.Println(err)
		}
		saveBasedOnStatusCode(myUrl, resp.StatusCode)
		parseOutput(resp)
		listOfProcessedUrls = append(listOfProcessedUrls, myUrl)
	}

	fmt.Println(listOfProcessedUrls)
	fmt.Println(listOfUrls)

	if len(listOfProcessedUrls) != len(listOfUrls) {
		collectURLOutput()
	}
}

func parseOutput(resp *http.Response) []string {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	processAllProduct(doc)
	return []string{"test"}
}

func processAllProduct(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		// process the Product details within each <a> element
		processNode(n)

	}
	// traverse the child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		processAllProduct(c)
	}
}

func processNode(n *html.Node) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			// retrieve src value
			ahref := a.Val
			// print image UR
			newUrl := ahref
			if !strings.Contains(ahref, "http") {
				newUrl = fmt.Sprintf("%s/%s", url, ahref)
			}
			if !slices.Contains(listOfUrls, newUrl) && strings.Contains(newUrl, url) {
				listOfUrls = append(listOfUrls, newUrl)
			}
			//fmt.Println("[URL]: ", newUrl)
		}
	}
}

func main() {
	url = "https://scrape-me.dreamsofcode.io"
	listOfUrls = append(listOfUrls, url)

	collectURLOutput()
	fmt.Println(listOfUrls)

	fmt.Println(sortedResults)

}
