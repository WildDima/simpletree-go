package woody

import (
	"reflect"
	"testing"
)

func TestAddChildren(t *testing.T) {
	base := &Node{Value: "First"}

	tests := []struct {
		input interface{}
		want  Node
	}{
		{input: "Hi", want: Node{Value: "Added"}},
	}

	for _, test := range tests {
		if got, err := base.AddChildren(test.input); reflect.DeepEqual(*got, test.want) || err != nil {
			t.Errorf("AddChildren(%v) = %v", test.input, got)
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
	base.AddChildren(&Node{Value: "Will remove it"})

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

func BenchmarkAddChildren(b *testing.B) {
	base := &Node{Value: "First"}

	for n := 0; n < b.N; n++ {
		base.AddChildren(1)
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
	base.AddChildren(&Node{Value: "Remove"})

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
