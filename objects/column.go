package objects

import (
	"errors"
	"github.com/antchfx/xmlquery"
)

type Column struct {
	Id            string `json:"-"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Comment       string `json:"comment"`
	DefaultValue  string `json:"default_value"`
	Nullable      bool   `json:"nullable"`
	AutoIncrement bool   `json:"auto_increment"`
	Unsigned      bool   `json:"unsigned"`
	Length        int    `json:"length"`
	Precision     int    `json:"precision"`
	Scale         int    `json:"scale"`
}

func parseColumns(tableNode *xmlquery.Node) ([]Column, error) {
	columnNodes := xmlquery.Find(tableNode, "value[@key=\"columns\"]/value[@struct-name=\"db.mysql.Column\"]")
	var results []Column
	for _, columnNode := range columnNodes {
		column, err := parseColumn(columnNode)
		if err == nil {
			results = append(results, column)
		}
	}

	return results, nil
}

func parseColumn(columnNode *xmlquery.Node) (Column, error) {
	column := Column{
		Id:            "",
		Name:          "",
		Comment:       "",
		Type:          "",
		DefaultValue:  "",
		Nullable:      false,
		AutoIncrement: false,
		Unsigned:      false,
		Length:        0,
		Precision:     0,
		Scale:         0,
	}

	column.Id = columnNode.SelectAttr("id")
	name, err := getValue(columnNode, "name")
	if err == nil {
		column.Name = name
	}
	comment, err := getValue(columnNode, "comment")
	if err == nil {
		column.Comment = comment
	}
	defaultValue, err := getValue(columnNode, "defaultValue")
	if err == nil {
		column.DefaultValue = defaultValue
	}
	nullable, err := getValueBool(columnNode, "isNotNull")
	if err == nil {
		column.Nullable = nullable
	}
	autoIncrement, err := getValueBool(columnNode, "autoIncrement")
	if err == nil {
		column.AutoIncrement = autoIncrement
	}

	flags, err := getList(columnNode, "flags")
	for _, flag := range flags {
		switch flag {
		case "UNSIGNED":
			column.Unsigned = true
		}
	}

	length, err := getValueInt(columnNode, "length")
	if err == nil {
		column.Length = length
	}
	precision, err := getValueInt(columnNode, "precision")
	if err == nil {
		column.Precision = precision
	}
	scale, err := getValueInt(columnNode, "scale")
	if err == nil {
		column.Scale = scale
	}

	simpleType, err := getLink(columnNode, "simpleType")
	if err == nil {
		column.Type = getDotNotationLastElement(simpleType)
	}

	if column.Id == "" {
		return column, errors.New("cannot parse table")
	}

	return column, nil
}
