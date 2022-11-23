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
	// URL
	path := "https://www.fos.kuis.kyoto-u.ac.jp/~igarashi/CoPL/index.cgi"
	u, err := url.Parse(path)
	if err != nil {
		panic(err)
	}
	// session cookie
	loginName := os.Getenv("loginName")
	c := http.DefaultClient
	cookie := http.Cookie{Name: "loginas", Value: loginName}
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	c.Jar = jar
	c.Jar.SetCookies(u, []*http.Cookie{&cookie})
	// request body
	v := url.Values{}
	ans := `
	S(S(Z)) is less than S(S(S(Z))) by L-Succ{}
	`
	v.Add("derivation", ans)
	v.Add("command", "answer")
	v.Add("game", "CompareNat1")
	v.Add("no", "9")
	// request header
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(v.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://www.fos.kuis.kyoto-u.ac.jp/~igarashi/CoPL/index.cgi?qno=9")
	// post
	res, err := c.Do(req)
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
