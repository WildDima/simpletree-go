package simpletree

import (
	"testing"
)

func TestAddChildren(t *testing.T) {
	base := fillTree(10)

	tests := []struct {
		input interface{}
		want  Node
	}{
		{input: "1", want: Node{Value: "1"}},
		{input: 1, want: Node{Value: 1}},
		{input: struct {
			a string
			b int
		}{a: "Hi", b: 123},
			want: Node{
				Value: struct {
					a string
					b int
				}{
					a: "Hi",
					b: 123,
				},
			},
		},
	}

	for _, test := range tests {
		if got, err := base.AddChildren(test.input); got.Value != test.want.Value || err != nil {
			t.Errorf("AddChildren(%v) = %v", test.input, got.Value)
		}
	}
}

func TestSetChildren(t *testing.T) {
	base := &Node{Value: "First"}

	tests := []struct {
		input interface{}
		want  Node
	}{
		{input: "Hi", want: Node{Value: "Hi"}},
	}

	for _, test := range tests {
		if got, err := base.setChildren(test.input); got.Value.(string) != test.want.Value.(string) || err != nil {
			t.Errorf("setChildren(%v) = %v", test.input, got.Value)
		}
	}
}

func TestAddSibling(t *testing.T) {
	base := &Node{Value: "First"}

	tests := []struct {
		input interface{}
		want  Node
	}{
		{input: "Hi", want: Node{Value: "Hi"}},
	}

	for _, test := range tests {
		if got, err := base.addSibling(test.input); got.Value.(string) != test.want.Value.(string) || err != nil {
			t.Errorf("addSibling(%v) = %v", test.input, got.Value)
		}
	}
}

func TestRemoveChildren(t *testing.T) {
	base := &Node{Value: "First"}
	base.setChildren(&Node{Value: "Will remove it"})

	if got, err := base.removeChildren(); got == nil || base.Children != nil || err != nil {
		t.Errorf("removeChildren() should remove children, but wasn't: %v", base)
	}
}

func TestRemoveSibling(t *testing.T) {
	base := &Node{Value: "First"}
	base.addSibling(&Node{Value: "Will remove it"})

	if got, err := base.removeSibling(); got == nil || base.Children != nil || err != nil {
		t.Errorf("removeSibling() should remove sibling, but wasn't: %v", base)
	}
}

func TestFind(t *testing.T) {
	rootNode := fillTree(10)

	lambda := func(n *Node) (b bool) {
		b = false

		if n.Value.(int) == 2 {
			b = true
		}

		return
	}

	if got, _ := rootNode.Find(lambda); got == nil {
		t.Errorf("Find() should return node, but was: %v, got")
	}
}

func TestDeleteIf(t *testing.T) {
	rootNode := fillTree(10)

	lambda := func(n *Node) (b bool) {
		b = false

		if n.Value.(int) == 2 {
			b = true
		}

		return
	}

	if got, _ := rootNode.DeleteIf(lambda); got == nil {
		t.Errorf("Delete() should return node, but was: %v, got")
	}

	rootNode = fillTree(5)

	lambda = func(n *Node) (b bool) {
		b = false

		if n.Value.(int) == 2 {
			b = true
		}

		return
	}

	rootNode.DeleteIf(lambda)
	//if got, _ := rootNode.DeleteIf(lambda); got != nil {
	if _, res := rootNode.DeleteIf(lambda); !res {
		t.Errorf("Delete() should return false, but was: %v", res)
	}
}

func TestSelect(t *testing.T) {
	rootNode := fillTree(5)

	lambda := func(n *Node) (b bool) {
		b = false

		if n.Value.(int) == 4 {
			b = true
		}

		return
	}

	if got, _ := rootNode.Select(lambda); len(got) != 4 {
		t.Errorf("Find() should return 4 nodes, but was: %v", len(got))
	}
}

func TestSize(t *testing.T) {
	base := fillTree(5)

	if got := base.Size(); got != 21 {
		t.Errorf("Size() = %v", got)
	}

	base = &Node{Value: 0}
	base.AddChildren(1)
	base.AddChildren(2)
	base.AddChildren(3)

	if got := base.Size(); got != 4 {
		t.Errorf("Size() = %v", got)
	}
}

func TestNext(t *testing.T) {
	size := 5
	base := fillTree(size)

	iterator := base.NewDeepFirstSearch()

	var got *Node

	for n := 1; n < size*(size-1)+2; n++ {
		got, _ = iterator.Next()
	}
	if got.Parent != nil {
		t.Errorf("iterator.Next() != %v, but == %v", got.Value)
	}
}

func BenchmarkAddChildren(b *testing.B) {
	base := &Node{Value: "First"}

	for n := 0; n < b.N; n++ {
		base.AddChildren(n)
	}
}

func BenchmarkSetChildren(b *testing.B) {
	base := &Node{Value: "First"}

	for n := 0; n < b.N; n++ {
		base.setChildren(1)
	}
}

func BenchmarkAddSibling(b *testing.B) {
	base := &Node{Value: "First"}

	for n := 0; n < b.N; n++ {
		base.addSibling(1)
	}
}

func BenchmarkRemoveChildren(b *testing.B) {
	base := &Node{Value: "First"}
	base.setChildren(&Node{Value: "Remove"})

	for n := 0; n < b.N; n++ {
		base.removeChildren()
	}
}

func BenchmarkRemoveSibling(b *testing.B) {
	base := &Node{Value: "First"}
	base.addSibling(&Node{Value: "Remove"})

	for n := 0; n < b.N; n++ {
		base.removeSibling()
	}
}

func BenchmarkNext(b *testing.B) {
	base := fillTree(50)
	iterator := base.NewDeepFirstSearch()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		iterator.Next()
	}
}

func BenchmarkFind(b *testing.B) {
	rootNode := fillTree(5)

	lambda := func(n *Node) (b bool) {
		b = false

		if n.Value.(int) == 2 {
			b = true
		}

		return
	}

	for n := 0; n < b.N; n++ {
		rootNode.Find(lambda)
	}
}

func BenchmarkSelect(b *testing.B) {
	rootNode := fillTree(5)

	lambda := func(n *Node) (b bool) {
		b = false

		if n.Value.(int) == 2 {
			b = true
		}

		return
	}

	for n := 0; n < b.N; n++ {
		rootNode.Select(lambda)
	}
}

func BenchmarkDeleteIf(b *testing.B) {
	rootNode := fillTree(5)

	lambda := func(n *Node) (b bool) {
		b = false

		if n.Value.(int) == 2 {
			b = true
		}

		return
	}

	for n := 0; n < b.N; n++ {
		rootNode.DeleteIf(lambda)
	}
}

func BenchmarkSize(b *testing.B) {
	rootNode := fillTree(10)

	for n := 0; n < b.N; n++ {
		rootNode.Size()
	}
}

func fillTree(depth int) *Node {
	nd := &Node{Value: 0}

	for i := 1; i < depth; i++ {
		child, _ := nd.AddChildren(i)

		for j := 1; j < depth; j++ {
			child.AddChildren(i + j)
		}
	}

	return nd
}
