package bst

import (
	"math/rand"
	"testing"
)

type testData struct {
	key, value int
}

func (d *testData) Compare(other Comparer) int {
	return d.key - other.(*testData).key
}

func verifyBst(tree *BST, t *testing.T) {
	n := tree.First()
	if n == nil {
		return
	}
	lastValue := n.Value
	n = n.Successor()
	for ; n != nil; n = n.Successor() {
		if lastValue.Compare(n.Value) >= 0 {
			t.Errorf("BST property broken. Last value %d, current %d",
				lastValue.(*testData).key, n.Value.(*testData).key)
		}
	}
}

func verifyInsert(t *testing.T, tree *BST, key int) {
	searchVal := &testData{key, 0}
	tree.Insert(&testData{key, key})
	verifyBst(tree, t)
	// Get the node and check it contains correct value
	v := tree.Get(searchVal)
	if v.Value.(*testData).value != key {
		t.Errorf("Get Initial Value: For key %d, expected %d, got %d",
			key, key, v.Value.(*testData).value)
	}
	// Update the value and verify
	tree.Insert(&testData{key, key * 2})
	verifyBst(tree, t)
	v = tree.Get(searchVal)
	if v.Value.(*testData).value != key*2 {
		t.Errorf("Get Initial Value: For key %d, expected %d, got %d",
			key, key*2, v.Value.(*testData).value)
	}
}

func TestInsert(t *testing.T) {
	tree := New()
	values := [...]int{100, 50, 200, 75, 25, 150, 250}
	for i := 0; i < len(values); i++ {
		verifyInsert(t, tree, values[i])
	}
	// Test with random values
	for i := 0; i < 1024; i++ {
		verifyInsert(t, tree, rand.Int())
	}
}
