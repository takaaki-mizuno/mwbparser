package mwbparser

import (
	"testing"
)

func TestParse(t *testing.T) {
	tables, err := Parse("./data/db.mwb")
	if err != nil {
		t.Fatalf("failed to parse %#v", err)
	}
	if len(tables) < 1 {
		t.Fatal("failed to extract tables")
	}
}
