package CoPL

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHTMLResult(t *testing.T) {
	tests := [...]struct {
		caseName string
		testdata string
	}{
		{
			caseName: "hoge",
			testdata: "../testdata/html/test.html",
		},
	}
	for _, tt := range tests {
		t.Run(tt.caseName, func(t *testing.T) {
			f, err := os.Open(tt.testdata)
			assert.NoError(t, err)
			_, err = ParseHTMLResult(f)
			assert.NoError(t, err)
		})
	}
}
