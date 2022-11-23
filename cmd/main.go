package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	copl "github.com/kosuke1012/CoPL/CoPL"
)

func main() {
	loginName := os.Getenv("loginName")
	ans := `
	S(S(Z)) is less than S(S(S(Z))) by L-Succ{}
	`
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
