package footballoutsiders

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Traverse(parent *html.Node, nodeType atom.Atom, possibleClasses []string) []*html.Node {
	var matches []*html.Node
	child := parent.FirstChild

	for child != nil {

		// Search this node's attributes
		for _, v := range child.Attr {
			if child.DataAtom == atom.Table && v.Key == "class" && stringInSlice(strings.TrimSpace(v.Val), possibleClasses) {
				fmt.Println("Match: " + v.Key + "=" + v.Val + " : " + child.DataAtom.String())

				// Found a match
				matches = append(matches, child)
			}
		}

		// If this node has children, search that node recursively
		if child.FirstChild != nil {
			matches = append(matches, Traverse(child, nodeType, possibleClasses)...)
		}

		// Next Child
		child = child.NextSibling
	}

	return matches
}

type Table struct {
	node *html.Node
}

func (t *Table) Headers() ([]string, error) {
	var thead *html.Node

	// Find thead
	child := t.node.FirstChild
	for child != nil {
		if child.DataAtom == atom.Thead {
			thead = child
			break
		}

		child = child.NextSibling
	}

	// Find tr
	child = thead.FirstChild
	for child != nil {
		if child.DataAtom == atom.Tr {
			thead = child
			break
		}

		child = child.NextSibling
	}
	if child == nil {
		return nil, fmt.Errorf("unable to find table header row")
	}

	// Iterate through td
	return RowData(child), nil
}

func (t *Table) Data() ([][]string, error) {
	var data [][]string

	var tbody *html.Node

	// Find tbody
	child := t.node.FirstChild
	for child != nil {
		if child.DataAtom == atom.Tbody {
			tbody = child
			break
		}

		child = child.NextSibling
	}
	if child == nil {
		return nil, fmt.Errorf("unable to find table body")
	}

	// iterate through rows
	tr := tbody.FirstChild
	for tr != nil {
		if tr.DataAtom == atom.Tr {
			row := RowData(tr)
			data = append(data, row)
		}
		tr = tr.NextSibling
	}
	return data, nil
}

func (t *Table) Col(i int) ([]string, error) {
	data, err := t.Data()
	if err != nil {
		return nil, err
	}

	vals := make([]string, len(data))
	for y, v := range data {
		vals[y] = v[i]
	}
	return vals, nil
}

func (t *Table) Pretty() string {
	headers, err := t.Headers()
	if err != nil {
		panic(err)
	}

	data, err := t.Data()
	if err != nil {
		panic(err)
	}

	maxColLen := make([]int, len(headers))

	// Form header
	s := "| "
	for i, _ := range headers {
		col, err := t.Col(i)
		if err != nil {
			panic(err)
		}
		maxColLen[i] = maxLen(headers[i], col...)
		s = s + fmt.Sprintf("%"+strconv.Itoa(maxColLen[i])+"s", headers[i])
		if i+1 < len(headers) {
			s = s + " | "
		}
	}
	s = s + " |\n"

	fmt.Println(s)

	// Fill body

	for _, row := range data {
		s = s + "| "

		for i, cell := range row {
			fmt.Printf("index: %d\n", i)
			s = s + fmt.Sprintf("%"+strconv.Itoa(maxColLen[i])+"s", cell)
			if i+1 < len(row) {
				s = s + " | "
			}
		}

		s = s + " |\n"
	}

	return s
}

func RowData(tr *html.Node) []string {
	var cells []string
	// Iterate through td
	td := tr.FirstChild
	for td != nil {
		if td.DataAtom == atom.Td {
			cells = append(cells, RenderText(td))
		}

		td = td.NextSibling
	}
	return cells
}

func RenderText(node *html.Node) string {
	s := ""

	if node.DataAtom.String() == "" {
		s = s + " " + node.Data
	}

	child := node.FirstChild

	for child != nil {
		tmp := strings.TrimSpace(RenderText(child))
		if tmp != "" {
			s = s + " " + tmp
		}

		child = child.NextSibling
	}

	return strings.TrimSpace(s)
}

func maxLen(min string, s ...string) int {
	max := len(min)
	for _, v := range s {
		if len(v) > max {
			max = len(v)
		}
	}
	return max
}

func stringInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if s == v {
			return true
		}
	}
	return false
}
