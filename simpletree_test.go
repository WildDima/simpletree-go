package simpletree

import (
	"reflect"
	"testing"
)

func TestAddChildren(t *testing.T) {
	base := &Node{Value: "First"}
	for i := 0; i < 10; i++ {
		base.AddChildren(i)
	}

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

func TestsetChildren(t *testing.T) {
	base := &Node{Value: "First"}

	tests := []struct {
		input interface{}
		want  Node
	}{
		{input: "Hi", want: Node{Value: "Added"}},
	}

	for _, test := range tests {
		if got, err := base.setChildren(test.input); reflect.DeepEqual(*got, test.want) || err != nil {
			t.Errorf("setChildren(%v) = %v", test.input, got)
		}
	}
}

func TestAddSibling(t *testing.T) {
	base := &Node{Value: "First"}

	tests := []struct {
		input interface{}
		want  Node
	}{
		{input: "Hi", want: Node{Value: "Added"}},
	}

	for _, test := range tests {
		if got, err := base.AddSibling(test.input); reflect.DeepEqual(*got, test.want) || err != nil {
			t.Errorf("AddSibling(%v) = %v", test.input, got)
		}
	}
}

func TestRemoveChildren(t *testing.T) {
	base := &Node{Value: "First"}
	base.setChildren(&Node{Value: "Will remove it"})

	if got, err := base.RemoveChildren(); got == nil || base.Children != nil || err != nil {
		t.Errorf("RemoveChildren() should remove children, but wasn't: %v", base)
	}
}

func TestRemoveSibling(t *testing.T) {
	base := &Node{Value: "First"}
	base.AddSibling(&Node{Value: "Will remove it"})

	if got, err := base.RemoveSibling(); got == nil || base.Children != nil || err != nil {
		t.Errorf("RemoveSibling() should remove sibling, but wasn't: %v", base)
	}
}

func TestFind(t *testing.T) {
	rootNode := &Node{Value: 0}
	rootNode.AddChildren(1)
	anotherNode, _ := rootNode.AddChildren(3)
	anotherNode.AddChildren(2)

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

func TestNext(t *testing.T) {
	base := &Node{Value: "First"}
	base.AddChildren(&Node{Value: 1})
	base.AddChildren(&Node{Value: 2})
	n, _ := base.AddChildren(&Node{Value: 3})
	n.AddChildren(&Node{Value: 4})
	n.AddChildren(&Node{Value: 5})

	iterator := base.NewDeepFirstSearch()

	for n := 1; n < 6; n++ {
		if got, _ := iterator.Next(); got.Value == n {
			t.Errorf("iterator.Next() != %v", n)
		}
	}
}

func BenchmarkAddChildren(b *testing.B) {
	base := &Node{Value: "First"}

	for n := 0; n < b.N; n++ {
		base.AddChildren(n)
	}
}

func BenchmarksetChildren(b *testing.B) {
	base := &Node{Value: "First"}

	for n := 0; n < b.N; n++ {
		base.setChildren(1)
	}
}

func BenchmarkAddSibling(b *testing.B) {
	base := &Node{Value: "First"}

	for n := 0; n < b.N; n++ {
		base.AddSibling(1)
	}
}

func BenchmarkRemoveChildren(b *testing.B) {
	base := &Node{Value: "First"}
	base.setChildren(&Node{Value: "Remove"})

	for n := 0; n < b.N; n++ {
		base.RemoveChildren()
	}
}

func BenchmarkRemoveSibling(b *testing.B) {
	base := &Node{Value: "First"}
	base.AddSibling(&Node{Value: "Remove"})

	for n := 0; n < b.N; n++ {
		base.RemoveSibling()
	}
}

func BenchmarkNext(b *testing.B) {
	base := &Node{Value: "First"}
	iterator := base.NewDeepFirstSearch()
	for n := 0; n < 50; n++ {
		base.AddChildren(&Node{Value: n})
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		iterator.Next()
	}
}

func BenchmarkFind(b *testing.B) {
	rootNode := &Node{Value: 0}
	rootNode.AddChildren(1)
	anotherNode, _ := rootNode.AddChildren(3)
	anotherNode.AddChildren(2)

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

func fillTree(depth int) *Node {
	nd := &Node{Value: 0}

	for i := 1; i < depth; i++ {
		child, _ := nd.AddChildren(i)

		for j := 1; j < depth; j++ {
			child.AddChildren(j)
		}
	}

	return nd
}
