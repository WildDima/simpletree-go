package woody

import (
	"reflect"
	"testing"
)

func TestAddChildren(t *testing.T) {
	base := &Node{Value: "First"}

	tests := []struct {
		input interface{}
		obj   *Node
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

func BenchmarkAddChildren(b *testing.B) {
	base := &Node{Value: "First"}

	for n := 0; n < b.N; n++ {
		base.AddChildren(1)
	}
}
