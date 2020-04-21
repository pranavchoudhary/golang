package bst

///////////////////////////////////////////////////////////////////////////////////
type Comparer interface {
	// Types should implement a Compare method which takes one argument - the other
	// value to compare against. It returns negative value if "this" object is
	// smaller and a positive value if the "other" object is smaller. If, both are
	// same, zero is returned
	Compare(other Comparer) int
}

///////////////////////////////////////////////////////////////////////////////////
// Binary search tree node and methods
type Node struct {
	parent, right, left *Node
	Value               Comparer
}

func (n *Node) insert(value Comparer) {
	// TODO: If n == nil, then throw an exception
	cmpResult := n.Value.Compare(value)
	switch {
	case cmpResult > 0:
		if n.left == nil {
			n.left = &Node{Value: value, parent: n}
		} else {
			n.left.insert(value)
		}
	case cmpResult < 0:
		if n.right == nil {
			n.right = &Node{Value: value, parent: n}
		} else {
			n.right.insert(value)
		}
	case cmpResult == 0:
		n.Value = value
	}
}

func (n *Node) get(value Comparer) *Node {
	if n == nil {
		return nil
	}
	cmpResult := n.Value.Compare(value)
	switch {
	case cmpResult > 0:
		return n.left.get(value)
	case cmpResult < 0:
		return n.right.get(value)
	default:
		return n
	}
}

func (n *Node) first() *Node {
	for ; n != nil && n.left != nil; n = n.left {
	}
	return n
}

func (n *Node) Successor() *Node {
	// If there is a right subtree, return the smaller in right subtree
	if n.right != nil {
		return n.right.first()
	}
	// Find the first left child in the ancestors
	for {
		if n.parent == nil {
			return nil
		} else if n.parent.left == n {
			return n.parent
		} else {
			n = n.parent
		}
	}
	return n
}

///////////////////////////////////////////////////////////////////////////////////
// BST and methods
type BST struct {
	head *Node
}

// Create a new BST instance and return its pointer
func New() *BST {
	return &BST{}
}

// Insert a new node in the BST
func (t *BST) Insert(value Comparer) {
	if t.head == nil {
		t.head = &Node{Value: value}
		return
	}
	t.head.insert(value)
}

// Find a node in the BST
func (t *BST) Get(value Comparer) *Node {
	return t.head.get(value)
}

// Return the first node (smallest by value)
func (t *BST) First() *Node {
	return t.head.first()
}
