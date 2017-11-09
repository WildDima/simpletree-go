package simpletree

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

type lambda func(n *Node) bool

func NewNode(i interface{}) *Node {
	n := &Node{Value: i}
	return n
}

func (c *Node) Find(l lambda) (n *Node, res bool) {
	dfs := c.NewDeepFirstSearch()
	for {
		n, res = dfs.Next()
		if !res {
			n, res = nil, false
			return
		}

		if l(n) {
			res = true
			return
		}
	}

	n, res = nil, false
	return
}

func (c *Node) AddChildren(i interface{}) (n *Node, err error) {
	if c.Children == nil {
		c.Children = &Node{Value: i, Parent: c}
		return c.Children, nil
	}

	child := c.Children

	for {
		if child.Sibling == nil {
			n = &Node{Value: i, Parent: c}
			child.Sibling = n
			break
		}

		child = child.Sibling
	}

	return
}

//func (c *Node) DeleteIf(l lambda) (n *Node, res bool) {
//n = c.Find(l)

//n.Delete()
//}

func (c *Node) Select(l lambda) (ns []*Node, res bool) {
	dfs := c.NewDeepFirstSearch()

	for {
		var n *Node

		n, res = dfs.Next()

		if !res {
			break
		}

		if l(n) {
			ns = append(ns, n)
		}
	}

	res = len(ns) == 0

	return
}

func (c *Node) NewDeepFirstSearch() (dfs *DeepFirstSearch) {
	dfs = new(DeepFirstSearch)
	dfs.currentNode = c
	dfs.stack = arraystack.New()
	dfs.stack.Push(c)
	dfs.visited = !c.visited

	return dfs
}

func (dfs *DeepFirstSearch) Next() (n *Node, res bool) {
	i, present := dfs.stack.Peek()
	res = true

	if !present {
		n, res = nil, false
		fmt.Println(n, res)
		return
	}

	p := i.(*Node)

	if p.Children != nil && p.Children.visited != dfs.visited {
		dfs.stack.Push(p.Children)
		dfs.currentNode = p.Children
		dfs.currentNode.visited = dfs.visited
		n = dfs.currentNode
	} else if p.Sibling != nil && p.Sibling.visited != dfs.visited {
		dfs.stack.Push(p.Sibling)
		dfs.currentNode = p.Sibling
		dfs.currentNode.visited = dfs.visited
		n = dfs.currentNode
	} else if p.Parent != nil {
		dfs.stack.Pop()
		dfs.currentNode = p.Parent
		dfs.currentNode.visited = dfs.visited
		n = dfs.currentNode
	} else {
		n, res = nil, false
	}

	return
}

func (c *Node) setChildren(i interface{}) (n *Node, err error) {
	if c.Children != nil {
		n, err = nil, fmt.Errorf("Children node for %v alreade exist", c)
	}

	c.Children = &Node{Value: i, Parent: c}

	n = c.Children

	return
}

func (c *Node) addSibling(i interface{}) (n *Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			n, err = nil, fmt.Errorf("Error while adding %q to node %q", n, c)
		}
	}()

	c.Sibling = &Node{Value: i, Parent: c.Parent}

	n = c.Sibling

	return
}

func (c *Node) removeChildren() (n *Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			n, err = nil, fmt.Errorf("Error while adding %q to node %q", n, c)
		}
	}()

	c.Children = nil

	n = c

	return
}

func (c *Node) removeSibling() (n *Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			n, err = nil, fmt.Errorf("Error while adding %q to node %q", n, c)
		}
	}()

	c.Sibling = nil

	n = c

	return
}
