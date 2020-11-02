package index

// Tree - binary search tree
// based on https://github.com/hardcode-dev/go-core/blob/master/5-ds/bst/bst.go
type Tree struct {
	root      *Element
	currentID uint
}

// Element - tree element
type Element struct {
	left, right *Element
	Value       *Document
}

// Document - stores information about web page
type Document struct {
	ID    uint
	URL   string
	Title string
}

// Insert document in Tree, autoincrements currentID, returns id of inserted document
func (t *Tree) Insert(doc *Document) uint {
	doc.ID = t.currentID
	t.currentID++

	e := &Element{Value: doc}
	if t.root == nil {
		t.root = e
		return doc.ID
	}

	return t.insert(t.root, e)
}

// insert recursive
func (t *Tree) insert(node, new *Element) uint {
	if new.Value.ID < node.Value.ID {
		if node.left == nil {
			node.left = new
			return new.Value.ID
		}

		return t.insert(node.left, new)
	}

	if node.right == nil {
		node.right = new
		return new.Value.ID
	}

	return t.insert(node.right, new)
}

// Search for document in tree
func (t *Tree) Search(id uint) *Document {
	return search(t.root, id)
}

// search recursive
func search(el *Element, id uint) *Document {
	if el == nil {
		return nil
	}

	if el.Value.ID == id {
		return el.Value
	}

	if el.Value.ID < id {
		return search(el.right, id)
	}

	return search(el.left, id)
}
