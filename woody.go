package woody

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
)

type Node struct {
	Value    interface{}
	Parent   *Node
	Children *Node
	Sibling  *Node
	visited  bool
}

type DeepFirstSearch struct {
	stack       *arraystack.Stack
	currentNode *Node
	visited     bool
}

func (c *Node) setChildren(i interface{}) (*Node, error) {
	if c.Children != nil {
		return nil, fmt.Errorf("Children node for %v alreade exist", c)
	}

	c.Children = &Node{Value: i, Parent: c}

	return c.Children, nil
}

func NewNode(i interface{}) *Node {
	n := &Node{Value: i}
	return n
}

func (c *Node) AddChildren(i interface{}) (*Node, error) {
	var n *Node
	if c.Children == nil {
		c.Children = &Node{Value: i, Parent: c}
		return c.Children, nil
	}

	child := c.Children

	for {
		if child.Sibling == nil {
			n = &Node{Value: i, Parent: c}
			break
		}

		child = child.Sibling
	}

	return n, nil
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

func (c *Node) NewDeepFirstSearch() (dfs *DeepFirstSearch) {
	dfs = new(DeepFirstSearch)
	dfs.currentNode = c
	dfs.stack = arraystack.New()
	dfs.stack.Push(c)
	dfs.visited = !c.visited

	return dfs
}

func (dfs *DeepFirstSearch) Next() (n *Node, err error) {
	i, present := dfs.stack.Peek()

	if !present {
		n = nil
		err = fmt.Errorf("Finished!")
		return
	}

	p := i.(*Node)

	if p.Children != nil && p.Children.visited != dfs.visited {
		dfs.stack.Push(p.Children)
		dfs.currentNode = p.Children
		n = dfs.currentNode
	} else if p.Sibling != nil && p.Sibling.visited != dfs.visited {
		dfs.stack.Push(p.Sibling)
		dfs.currentNode = p.Sibling
		n = dfs.currentNode
	} else if p.Parent != nil {
		dfs.stack.Pop()
		dfs.currentNode = p.Parent
		n = dfs.currentNode
	} else {
		n, err = nil, fmt.Errorf("Finished!")
	}

	return
}
