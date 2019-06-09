package objects

import (
	"errors"
	"github.com/antchfx/xmlquery"
	"strconv"
	"strings"
)

func getValue(node *xmlquery.Node, name string) (string, error) {
	resultNode := xmlquery.FindOne(node, "value[@key=\""+name+"\"]")
	if resultNode == nil {
		return "", errors.New("cannot find key: " + name)
	}
	return resultNode.InnerText(), nil
}

func getValueInt(node *xmlquery.Node, name string) (int, error) {
	stringValue, err := getValue(node, name)
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(stringValue)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func getValueBool(node *xmlquery.Node, name string) (bool, error) {
	stringValue, err := getValue(node, name)
	if err != nil {
		return false, err
	}
	if stringValue == "1" {
		return true, nil
	}

	return false, nil
}

func getLink(node *xmlquery.Node, name string) (string, error) {
	resultNode := xmlquery.FindOne(node, "link[@key=\""+name+"\"]")
	if resultNode == nil {
		return "", errors.New("cannot find key: " + name)
	}
	return resultNode.InnerText(), nil
}

func getDotNotationLastElement(text string) string {
	elements := strings.Split(text, ".")
	return elements[len(elements)-1]
}

func getList(node *xmlquery.Node, name string) ([]string, error) {
	parentNode := xmlquery.FindOne(node, "value[@key=\""+name+"\"]")
	if parentNode == nil {
		return nil, errors.New("cannot find key: " + name)
	}
	childNodes := xmlquery.Find(parentNode, "value")
	var result []string
	for _, columnNode := range childNodes {
		result = append(result, columnNode.InnerText())
	}

	return result, nil
}
