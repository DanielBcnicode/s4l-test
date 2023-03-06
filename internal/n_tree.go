package internal

// In this file I've implemented a N-TREE structure and algorithm to achieve
// the s4l maximize booking problem. The current implementation is coupled to
// the solution. It can be refactored to have a generic N-TREE implementation
// to use any kind of data inside.
// The solution proposed is a double-linked N-TREE

// NodeType is the data type used to indicate the node type: Node, Leaf, Origin
type NodeType string

const (
	NODE   = NodeType("Node")
	LEAF   = NodeType("Leaf")
	ORIGIN = NodeType("Origin")
)

// Node is the basic element to construct a N-TREE
type Node struct {
	children       []*Node
	father         *Node
	kind           NodeType
	profitPerNight float32
	profit         float32
	innerData      interface{}
}

// NewNode is the Node constructor. If father is null we assume it is the ROOT
func NewNode(father *Node) *Node {
	kind := LEAF
	if father == nil {
		kind = ORIGIN
	}

	node := &Node{
		children:  nil,
		father:    father,
		kind:      kind,
		innerData: nil,
	}

	if father != nil {
		father.AddChildren(node)
	}

	return node
}

// AddChildren inserts a children node in the current one
func (n *Node) AddChildren(child *Node) {
	child.father = n
	if child.ChildrenNumber() == 0 {
		child.kind = LEAF
	} else {
		child.kind = NODE
	}

	n.children = append(n.children, child)
	if n.kind == LEAF {
		n.kind = NODE
	}
}

// Children returns the array of children (Nodes) for the current node
func (n *Node) Children() []*Node {
	return n.children
}

// SetData sets the generic data from node. I use innerData to store the booking ID
func (n *Node) SetData(data interface{}) {
	n.innerData = data
}

// Data returns the data stored in innerData
func (n *Node) Data() interface{} {
	return n.innerData
}

// ChildrenNumber returns the number of direct children
func (n *Node) ChildrenNumber() int {
	return len(n.children)
}

// NodeCount returns the number of dependencies nodes for the current one
func (n *Node) NodeCount() int {
	nn := 1
	for _, child := range n.children {
		nn += child.NodeCount()
	}

	return nn
}

// ProfitPerNight gets the profit per night value
func (n *Node) ProfitPerNight() float32 {
	return n.profitPerNight
}

// SetProfitPerNight sets the profit per night for the current Node
func (n *Node) SetProfitPerNight(profit float32) {
	n.profitPerNight = profit
}

// Profit returns the profit value for the current Node
func (n *Node) Profit() float32 {
	return n.profit
}

// SetProfit sets the profit for the current Node
func (n *Node) SetProfit(profit float32) {
	n.profit = profit
}

// Father returns the father node of the current one
func (n *Node) Father() *Node {
	return n.father
}

// GetLeafs returns all the dependencies leafs of the current node as an array of pointers
func (n *Node) GetLeafs() []*Node {
	var nn []*Node
	if n.kind == LEAF {
		nn = append(nn, n)
		return nn
	}

	for _, child := range n.children {
		nn = append(nn, child.GetLeafs()...)
	}

	return nn
}
