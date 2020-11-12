package html_parser_golang

//make use of /x/net/html lib https://godoc.org/golang.org/x/net/html

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

// Represents link <a> tag in html document
type Link struct {
	Href string
	Text string
}

func main() {
	fmt.Println("Welcome to html link parser: ")
	//Parse(doc)
}

// Takes in html doc prints in spaced format to show parent child relationships
func ParsePrintDoc(r io.Reader) {

	// get root of the Parse tree to traverse
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
    // do recursive search document for nodes
    dfsPrintDoc(doc, "")
}

// Takes in html doc prints in spaced format to show parent child relationships
func ParseGetLinks(r io.Reader) ([]Link, error) {

	// get root of the Parse tree to traverse
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	// do recursive search document for nodes
	dfsGetLinks(doc)
	return nil,nil
}

// Printing document in formatted way to see how related
func dfsPrintDoc(n *html.Node,padding string) {
	fmt.Println(padding, n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfsPrintDoc(c, padding+"  ")
	}
}

//
func dfsGetLinks(n *html.Node) []Link {
	links := make([]Link, 0)
	if n.Type == html.ElementNode && n.Data == "a" {
		newLink = Link{n.Href, n.Text}
		links = append(links, newLink)
	}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			return dfsGetLinks(c)
		}
	}
	return links
}



