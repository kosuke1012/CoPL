package CoPL

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func ParseHTMLResult(r io.Reader) (string, error) {
	t := html.NewTokenizer(r)
	for {

		switch tt := t.Next(); tt {
		case html.ErrorToken:
			return "", fmt.Errorf("div main id not found.")
		case html.StartTagToken:
			if _, hasAttr := t.TagName(); !hasAttr {
				continue
			}
			if key, val, _ := t.TagAttr(); string(key) == "id" && string(val) == "main" {
				var b strings.Builder
				for _, s := range []string{"h1", "pre"} {
					data, err := findStartTag(t, s)
					if err != nil {
						return "", err
					}
					fmt.Fprintln(&b, data)
				}
				return b.String(), nil
			}
		default:
			continue
		}
	}
}

func findStartTag(t *html.Tokenizer, tagName string) (string, error) {
	for {
		switch token := t.Token(); token.Type {
		case html.ErrorToken:
			return "", fmt.Errorf("%s tag not found.", tagName)
		case html.StartTagToken:
			if token.Data == tagName {
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
