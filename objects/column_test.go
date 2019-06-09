package objects

import (
	"github.com/antchfx/xmlquery"
	"os"
	"testing"
)

func TestParseColumns(t *testing.T) {
	reader, err := os.Open("../data/document.mwb.xml")
	if err != nil {
		t.Fatalf("failed to open xml file %#v", err)
	}
	document, err := xmlquery.Parse(reader)
	tableNode := xmlquery.FindOne(document, "//value[@struct-name=\"db.mysql.Table\"]")
	columns, err := parseColumns(tableNode)
	if err != nil {
		t.Fatalf("failed to parse columns %#v", err)
	}
	if len(columns) < 1 {
		t.Fatalf("failed to get columns %#v", err)
	}
}
