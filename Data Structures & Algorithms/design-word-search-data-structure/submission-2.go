
type DictNode struct {
	Children map[rune]*DictNode
	EOW      bool
}

func NewDictNode() *DictNode {
	return &DictNode{Children: make(map[rune]*DictNode)}
}

type WordDictionary struct {
	root *DictNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: NewDictNode()}
}

func (w *WordDictionary) AddWord(word string) {
	curr := w.root
	for _, w := range word {
		if node, ok := curr.Children[w]; ok {
			curr = node
			continue
		}
		node := NewDictNode()
		curr.Children[w] = node
		curr = node
	}
	curr.EOW = true
}

func (w *WordDictionary) Search(word string) bool {
	return search(w.root, []rune(word))
}

func search(root *DictNode, word []rune) bool {
	if len(word) == 0 {
		return root.EOW
	}

	r := word[0]
	node, ok := root.Children[r]
	if ok {
		return search(node, word[1:])
	}
	if r != '.' {
		return false
	}

	var result bool
	for _, ch := range root.Children {
		result = result || search(ch, word[1:])
	}

	return result
}

