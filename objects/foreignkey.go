package objects

import (
	"errors"
	"github.com/antchfx/xmlquery"
)

type ForeignKey struct {
	Id                 string   `json:"-"`
	Name               string   `json:"name"`
	Many               bool     `json:"many"`
	Columns            []string `json:"columns"`
	ColumnIds          []string `json:"-"`
	ReferenceColumns   []string `json:"referenceColumns"`
	ReferenceColumnIds []string `json:"-"`
	Comment            string   `json:"comment"`
	ReferenceTable     string   `json:"referenceTable"`
	ReferenceTableId   string   `json:"-"`
	DeleteRule         string   `json:"deleteRule"`
	UpdateRule         string   `json:"updateRule"`
}

func ParseForeignKeys(tableNode *xmlquery.Node, columns []Column, tables []Table) ([]ForeignKey, error) {
	foreignKeyNodes := xmlquery.Find(tableNode, "value[@key=\"foreignKeys\"]/value[@struct-name=\"db.mysql.ForeignKey\"]")
	var results []ForeignKey
	for _, foreignKeyNode := range foreignKeyNodes {
		foreignKey, err := parseForeignKey(foreignKeyNode, columns, tables)
		if err == nil {
			results = append(results, foreignKey)
		}
	}

	return results, nil
}

func parseForeignKey(foreignKeyNode *xmlquery.Node, columns []Column, tables []Table) (ForeignKey, error) {
	foreignKey := ForeignKey{
		Id:                 "",
		Name:               "",
		Many:               false,
		Columns:            []string{},
		ColumnIds:          []string{},
		ReferenceColumnIds: []string{},
		ReferenceColumns:   []string{},
		ReferenceTable:     "",
		ReferenceTableId:   "",
		DeleteRule:         "",
		UpdateRule:         "",
		Comment:            "",
	}
	foreignKey.Id = foreignKeyNode.SelectAttr("id")
	name, err := getValue(foreignKeyNode, "name")
	if err == nil {
		foreignKey.Name = name
	}
	comment, err := getValue(foreignKeyNode, "comment")
	if err == nil {
		foreignKey.Comment = comment
	}

	updateRule, err := getValue(foreignKeyNode, "updateRule")
	if err == nil {
		foreignKey.UpdateRule = updateRule
	}
	deleteRule, err := getValue(foreignKeyNode, "deleteRule")
	if err == nil {
		foreignKey.DeleteRule = deleteRule
	}

	referenceTableId, err := getLink(foreignKeyNode, "referencedTable")
	referenceTable := Table{
		Columns: []Column{},
	}
	if err == nil {
		for _, table := range tables {
			if table.Id == referenceTableId {
				referenceTable = table
				foreignKey.ReferenceTable = table.Name
				foreignKey.ReferenceTableId = referenceTableId
				break
			}
		}
	}

	many, err := getValueBool(foreignKeyNode, "Many")
	if err == nil {
		foreignKey.Many = many
	}

	columnNodes := xmlquery.Find(foreignKeyNode, ".//value[@key=\"columns\"]/link")
	for _, columnNode := range columnNodes {
		columnId := columnNode.InnerText()
		for _, column := range columns {
			if column.Id == columnId {
				foreignKey.Columns = append(foreignKey.Columns, column.Name)
				foreignKey.ColumnIds = append(foreignKey.ColumnIds, columnId)
				break
			}
		}

	}

	referenceColumnNodes := xmlquery.Find(foreignKeyNode, ".//value[@key=\"referencedColumns\"]/link")
	for _, columnNode := range referenceColumnNodes {
		columnId := columnNode.InnerText()
		for _, column := range referenceTable.Columns {
			if column.Id == columnId {
				foreignKey.ReferenceColumns = append(foreignKey.Columns, column.Name)
				foreignKey.ReferenceColumnIds = append(foreignKey.ColumnIds, columnId)
				break
			}
		}

	}

	if foreignKey.Id == "" {
		return foreignKey, errors.New("cannot parse table")
	}

	return foreignKey, nil
}
