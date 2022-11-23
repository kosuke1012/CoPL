package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/goccy/go-yaml"

	copl "github.com/kosuke1012/CoPL/CoPL"
)

var noFlag = flag.Int("n", 0, "Question no to answer.")

type workbook struct {
	Workbook []answer
}
type answer struct {
	No      string
	Game    string
	Anspath string
}

func main() {
	flag.Parse()

	var w workbook
	err := w.readWorkbook()
	if err != nil {
		panic(err)
	}
	var q answer
	for _, work := range w.Workbook {
		if work.No == strconv.Itoa(*noFlag) {
			q = work
			break
		}
	}
	if q.No == "" {
		panic("Question no not found.")
	}
	f, err := os.Open(q.Anspath)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	ans := string(b)
	loginName := os.Getenv("loginName")
	res, err := postAnswer(loginName, "9", "CompareNat1", ans)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// read response
	resStr, err := copl.ParseHTMLResult(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resStr)
}

func (w *workbook) readWorkbook() error {
	f, err := os.Open("../workbook.yaml")
	defer f.Close()
	if err != nil {
		return err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(b, w); err != nil {
		return err
	}
	return nil
}

func postAnswer(lname, q, game, ans string) (*http.Response, error) {
	// URL
	path := "https://www.fos.kuis.kyoto-u.ac.jp/~igarashi/CoPL/index.cgi"
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	// session cookie
	c := http.DefaultClient
	cookie := http.Cookie{Name: "loginas", Value: lname}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	c.Jar = jar
	c.Jar.SetCookies(u, []*http.Cookie{&cookie})
	// request body
	v := url.Values{}
	v.Add("derivation", ans)
	v.Add("command", "answer")
	v.Add("game", game)
	v.Add("no", q)
	// request header
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(v.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// post
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	return res, nil
}
