package d2tbl

import (
	"testing"
)

func exampleData() *TextDictionary {
	result := &TextDictionary{
		"abc":        "def",
		"someStr":    "Some long string",
		"teststring": "TeSt",
	}

	return result
}

func TestTBL_Marshal(t *testing.T) {
	tbl := exampleData()
	data := tbl.Marshal()

	newTbl, err := LoadTextDictionary(data)
	if err != nil {
		t.Error(err)
	}

	for key, value := range *tbl {
		newValue, ok := newTbl[key]

		if !ok {
			t.Fatal("string wasn't encoded to table")
		}

		if newValue != value {
			t.Fatal("unexpected value set")
		}
	}
}
