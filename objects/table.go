package objects

import (
	"errors"
	"github.com/antchfx/xmlquery"
)

type Table struct {
	Id          string         `json:"-"`
	Name        string         `json:"name"`
	Columns     []Column       `json:"columns"`
	Indexes     []Index        `json:"indexes"`
	ForeignKeys []ForeignKey   `json:"foreignKeys"`
	Node        *xmlquery.Node `json:"-"`
}

func ParseTables(documentNode *xmlquery.Node) ([]Table, error) {
	tableNodes := xmlquery.Find(documentNode, "//value[@struct-name=\"db.mysql.Table\"]")
	var tables []Table
	for _, tableNode := range tableNodes {
		table, err := parseTable(tableNode)
		if err == nil {
			tables = append(tables, table)
		}
	}

	var results []Table
	for _, table := range tables {
		foreignKeys, err := ParseForeignKeys(table.Node, table.Columns, tables)
		if err == nil {
			table.ForeignKeys = foreignKeys
			results = append(results, table)
		}
	}

	return results, nil
}

func parseTable(tableNode *xmlquery.Node) (Table, error) {
	table := Table{Id: ""}

	table.Id = tableNode.SelectAttr("id")
	name, err := getValue(tableNode, "name")
	if err == nil {
		table.Name = name
	}

	if table.Id == "" {
		return table, errors.New("cannot parse table")
	}

	table.Columns, err = parseColumns(tableNode)
	table.Indexes, err = parseIndexes(tableNode, table.Columns)

	table.Node = tableNode

	return table, nil
}
