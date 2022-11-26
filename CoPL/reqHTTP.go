package CoPL

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func PostAnswer(lname, q, game, ans string) (*http.Response, error) {
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
