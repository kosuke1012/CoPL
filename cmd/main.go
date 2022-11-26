package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

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
	res, err := copl.PostAnswer(loginName, "9", "CompareNat1", ans)
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
