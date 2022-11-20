package CoPL

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func ParseHTMLResult(r io.Reader) (string, error) {
	t := html.NewTokenizer(r)
	for {

		switch tt := t.Next(); tt {
		case html.ErrorToken:
			return "", fmt.Errorf("div main id not found.")
		case html.StartTagToken:
			if _, hasAttr := t.TagName(); hasAttr {
				if key, val, _ := t.TagAttr(); string(key) == "id" && string(val) == "main" {
					data, err := findH1Tag(t)
					if err != nil {
						panic(err)
					}
					return data, nil
				}
			}
			continue
		default:
			continue
		}
	}
}

func findH1Tag(t *html.Tokenizer) (string, error) {
	for {
		switch token := t.Token(); token.Type {
		case html.ErrorToken:
			return "", fmt.Errorf("h1 tag not found.")
		case html.StartTagToken:
			if token.Data == "h1" {
				t.Next()
				return t.Token().Data, nil
			}
			t.Next()
			continue
		default:
			t.Next()
			continue
		}
	}
}
