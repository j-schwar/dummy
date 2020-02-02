package dummy

import (
	"fmt"
	"testing"
)

func TestFixedNameGenerator_Count(t *testing.T) {
	tests := []struct {
		names []string
	}{
		{[]string{"a"}},
		{[]string{"a", "b"}},
		{[]string{"a", "b", "c"}},
		{[]string{"a", "b", "c", "d"}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Count_%d", len(test.names)), func(t *testing.T) {
			g := NewFixedNameGenerator(test.names...)
			if g.Count() != len(test.names) {
				t.Errorf("expected %d, got %d", len(test.names), g.Count())
			}
		})
	}
}

func TestFixedNameGenerator(t *testing.T) {
	set := []string{"Foo", "Bar", "Baz"}
	exists := func(x string) bool {
		for _, y := range set {
			if x == y {
				return true
			}
		}
		return false
	}

	g := NewFixedNameGenerator(set...)
	for range set {
		a, ok := g.Generate()
		if !ok || !exists(a) {
			t.Fatal("failed to generate expected output")
		}
	}

	_, ok := g.Generate()
	if ok {
		t.Fatal("expected generator to be exhausted")
	}
}
