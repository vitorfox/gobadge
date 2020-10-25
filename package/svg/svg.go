package svg

import (
	"fmt"
	"io"
	"strings"
)

type Svg Node

type NodeBuilder interface {
	Build(io.Writer, int, *Node)
}

type EmptyNodeBuilder struct{}

func (d *EmptyNodeBuilder) Build(o io.Writer, tab int, node *Node) {
	dataS := make([]string, 0)

	s := node

	for _, v := range s.Data {
		dataS = append(dataS, v.Build())
	}

	tabs := strings.Join(make([]string, tab), "  ")

	out := fmt.Sprintf("%s%s\n", tabs, strings.Join(dataS, " "))
	o.Write([]byte(out))
}

type DefaultNodeBuilder struct{}

func (d *DefaultNodeBuilder) Build(o io.Writer, tab int, node *Node) {
	dataS := make([]string, 0)

	s := node

	for _, v := range s.Data {
		dataS = append(dataS, v.Build())
	}

	tabs := strings.Join(make([]string, tab), "  ")

	end := "/>"

	if len(s.Child) > 0 {
		end = ">"
	}

	out := fmt.Sprintf("%s<%s %s%s\n", tabs, s.Tag, strings.Join(dataS, " "), end)
	o.Write([]byte(out))

	for _, child := range s.Child {
		child.build(o, tab+1)
	}

	if len(s.Child) > 0 {
		o.Write([]byte(fmt.Sprintf("%s</%s>\n", tabs, s.Tag)))
	}
}

func New(data ...Something) *Node {
	return &Node{
		Tag:     "svg",
		Child:   make([]*Node, 0),
		Data:    data,
		builder: new(DefaultNodeBuilder),
	}
}

func (s *Node) build(o io.Writer, tab int) {
	s.builder.Build(o, tab, s)
}

func (s *Node) Build(o io.Writer) {
	s.build(o, 1)
}

type Something interface {
	Build() string
}

type NumberElement struct {
	Name  string
	Value float64
}

func Number(name string, value float64) *NumberElement {
	return &NumberElement{
		Name:  name,
		Value: value,
	}
}

func (n NumberElement) Build() string {
	return fmt.Sprintf("%s=%d", n.Name, n.Value)
}

type StringElement struct {
	Name  string
	Value string
}

func String(name string, value string) *StringElement {
	return &StringElement{
		Name:  name,
		Value: value,
	}
}

func (n StringElement) Build() string {
	return fmt.Sprintf(`%s="%s"`, n.Name, n.Value)
}

type SimpleElement struct {
	Value string
}

func Simple(value string) *SimpleElement {
	return &SimpleElement{
		Value: value,
	}
}

func (n SimpleElement) Build() string {
	return n.Value
}

type Node struct {
	Child   []*Node
	Data    []Something
	Tag     string
	builder NodeBuilder
}

func (node *Node) Add(tag string, data ...Something) *Node {

	var bb NodeBuilder = new(DefaultNodeBuilder)

	if tag == "" {
		bb = new(EmptyNodeBuilder)
	}

	nNode := &Node{
		Tag:     tag,
		Data:    data,
		Child:   make([]*Node, 0),
		builder: bb,
	}

	node.Child = append(node.Child, nNode)

	return nNode
}
