package objects

import (
	"github.com/antchfx/xmlquery"
	"os"
	"testing"
)

func TestGetValue(t *testing.T) {
	reader, err := os.Open("../data/document.mwb.xml")
	if err != nil {
		t.Fatalf("failed to open xml file %#v", err)
	}
	document, err := xmlquery.Parse(reader)
	tableNode := xmlquery.FindOne(document, "//value[@struct-name=\"db.mysql.Table\"]")
	_, err = getValue(tableNode, "name")
	if err != nil {
		t.Fatalf("failed to use getValue %#v", err)
	}
}

func TestGetLink(t *testing.T) {
	reader, err := os.Open("../data/document.mwb.xml")
	if err != nil {
		t.Fatalf("failed to open xml file %#v", err)
	}
	document, err := xmlquery.Parse(reader)
	tableNode := xmlquery.FindOne(document, "//value[@struct-name=\"db.mysql.Table\"]")
	_, err = getLink(tableNode, "primaryKey")
	if err != nil {
		t.Fatalf("failed to use getLink %#v", err)
	}
}

func TestGetValueBool(t *testing.T) {
	reader, err := os.Open("../data/document.mwb.xml")
	if err != nil {
		t.Fatalf("failed to open xml file %#v", err)
	}
	document, err := xmlquery.Parse(reader)
	tableNode := xmlquery.FindOne(document, "//value[@struct-name=\"db.mysql.Table\"]")
	_, err = getValueBool(tableNode, "checksum")
	if err != nil {
		t.Fatalf("failed to use getValueBool %#v", err)
	}
}

func TestGetValueInt(t *testing.T) {
	reader, err := os.Open("../data/document.mwb.xml")
	if err != nil {
		t.Fatalf("failed to open xml file %#v", err)
	}
	document, err := xmlquery.Parse(reader)
	tableNode := xmlquery.FindOne(document, "//value[@struct-name=\"db.mysql.Table\"]")
	_, err = getValueInt(tableNode, "checksum")
	if err != nil {
		t.Fatalf("failed to use getValueInt %#v", err)
	}
}
