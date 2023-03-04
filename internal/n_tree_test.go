package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//		 Tree structure to test
//		             n0
//			       /    \
//			   n1a        n1b
//			/      \
//	   n2a1        n2a2
//		             /     \
//		          n3b1      n3b2
func TestNewNode(t *testing.T) {
	n0 := NewNode(nil)
	n1a := NewNode(n0)
	n1b := NewNode(n0)
	n1b.SetData(3.12)
	n2a1 := NewNode(n1a)
	n2a2 := NewNode(n1a)
	n2a2.SetData("data test")
	n3b1 := NewNode(n2a2)
	n3b2 := NewNode(nil)
	n2a2.AddChildren(n3b2)
	leafs := []*Node{n1b, n2a1, n3b1, n3b2}

	// Children number tests
	assert.Equal(t, 2, n0.ChildrenNumber())
	assert.Equal(t, 0, n1b.ChildrenNumber())
	assert.Equal(t, 0, n2a1.ChildrenNumber())
	assert.Equal(t, 2, n2a2.ChildrenNumber())
	assert.Equal(t, 0, n3b1.ChildrenNumber())
	assert.Equal(t, 0, n3b2.ChildrenNumber())

	// Node count test
	assert.Equal(t, 7, n0.NodeCount())
	assert.Equal(t, 5, n1a.NodeCount())
	assert.Equal(t, 1, n1b.NodeCount())
	assert.Equal(t, 1, n2a1.NodeCount())
	assert.Equal(t, 3, n2a2.NodeCount())
	assert.Equal(t, 1, n3b1.NodeCount())
	assert.Equal(t, 1, n3b2.NodeCount())

	// Leafs test
	assert.Equal(t, 4, len(n0.GetLeafs()))
	assert.ElementsMatch(t, leafs, n0.GetLeafs())

	// Data test
	assert.Equal(t, "data test", n2a2.Data())
	assert.Equal(t, 3.12, n1b.Data())
}
