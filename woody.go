package woody

import (
	"fmt"
)

type Tree struct {
	Nodes []*Node
}

type Node struct {
	Value    interface{}
	Parent   *Node
	Children *Node
	Sibling  *Node
}

func (c *Node) AddChildren(i interface{}) (n *Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			n, err = nil, fmt.Errorf("Error while adding %q to node %q", n, c)
		}
	}()

	c.Children = &Node{Value: i, Parent: c}

	n = c.Children

	return n, err
}

func (c *Node) AddSibling(i interface{}) (n *Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			n, err = nil, fmt.Errorf("Error while adding %q to node %q", n, c)
		}
	}()

	c.Sibling = &Node{Value: i, Parent: c.Parent}

	n = c.Sibling

	return n, nil
}

func (c *Node) RemoveChildren() (n *Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			n, err = nil, fmt.Errorf("Error while adding %q to node %q", n, c)
		}
	}()

	c.Children = nil

	n = c

	return n, nil
}

func (c *Node) RemoveSibling() (n *Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			n, err = nil, fmt.Errorf("Error while adding %q to node %q", n, c)
		}
	}()

	c.Sibling = nil

	n = c

	return n, nil
}

//func (c *Node) Next() (Node, error) {
//}

//func DFS(func(t *Tree)) {}
