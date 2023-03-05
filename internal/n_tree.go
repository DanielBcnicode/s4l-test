package internal

type NodeType string

const (
	NODE   = NodeType("Node")
	LEAF   = NodeType("Leaf")
	ORIGIN = NodeType("Origin")
)

type Node struct {
	children       []*Node
	father         *Node
	kind           NodeType
	profitPerNight float32
	profit         float32
	innerData      interface{}
}

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

func (n *Node) Children() []*Node {
	return n.children
}
func (n *Node) SetData(data interface{}) {
	n.innerData = data
}

func (n *Node) Data() interface{} {
	return n.innerData
}

func (n *Node) ChildrenNumber() int {
	return len(n.children)
}

func (n *Node) NodeCount() int {
	nn := 1
	for _, child := range n.children {
		nn += child.NodeCount()
	}

	return nn
}

func (n *Node) ProfitPerNight() float32 {
	return n.profitPerNight
}

func (n *Node) SetProfitPerNight(profit float32) {
	n.profitPerNight = profit
}

func (n *Node) Profit() float32 {
	return n.profit
}

func (n *Node) SetProfit(profit float32) {
	n.profit = profit
}

func (n *Node) Father() *Node {
	return n.father
}

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
