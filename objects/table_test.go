package objects

import (
	"github.com/antchfx/xmlquery"
	"os"
	"testing"
)

func TestParseTables(t *testing.T) {
	reader, err := os.Open("../data/document.mwb.xml")
	if err != nil {
		t.Fatalf("failed to open xml file %#v", err)
	}
	document, err := xmlquery.Parse(reader)
	tables, err := ParseTables(document)
	if err != nil {
		t.Fatalf("failed to parse tables %#v", err)
	}
	if len(tables) < 1 {
		t.Fatalf("failed to get tables %#v", err)
	}
}
