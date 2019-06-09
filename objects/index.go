package objects

import (
	"errors"
	"github.com/antchfx/xmlquery"
)

type Index struct {
	Id        string   `json:"-"`
	Name      string   `json:"name"`
	Primary   bool     `json:"primary"`
	Unique    bool     `json:"unique"`
	Columns   []string `json:"columns"`
	ColumnIds []string `json:"-"`
	Comment   string   `json:"comment"`
}

func parseIndexes(tableNode *xmlquery.Node, columns []Column) ([]Index, error) {
	indexNodes := xmlquery.Find(tableNode, "value[@key=\"indices\"]/value[@struct-name=\"db.mysql.Index\"]")
	var results []Index
	for _, indexNode := range indexNodes {
		index, err := parseIndex(indexNode, columns)
		if err == nil {
			results = append(results, index)
		}
	}

	return results, nil
}

func parseIndex(indexNode *xmlquery.Node, columns []Column) (Index, error) {
	index := Index{
		Id:        "",
		Name:      "",
		Primary:   false,
		Unique:    false,
		Columns:   []string{},
		ColumnIds: []string{},
		Comment:   "",
	}
	index.Id = indexNode.SelectAttr("id")
	name, err := getValue(indexNode, "name")
	if err == nil {
		index.Name = name
	}
	comment, err := getValue(indexNode, "comment")
	if err == nil {
		index.Comment = comment
	}

	isPrimary, err := getValueBool(indexNode, "isPrimary")
	if err == nil {
		index.Primary = isPrimary
	}

	unique, err := getValueBool(indexNode, "unique")
	if err == nil {
		index.Unique = unique
	}

	columnNodes := xmlquery.Find(indexNode, ".//value[@struct-name=\"db.mysql.IndexColumn\"]")
	for _, columnNode := range columnNodes {
		columnId, err := getLink(columnNode, "referencedColumn")
		if err == nil {
			for _, column := range columns {
				if column.Id == columnId {
					index.Columns = append(index.Columns, column.Name)
					index.ColumnIds = append(index.ColumnIds, columnId)
					break
				}
			}
		}
	}

	if index.Id == "" {
		return index, errors.New("cannot parse table")
	}

	return index, nil
}
