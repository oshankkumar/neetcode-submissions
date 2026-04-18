func Constructor() PrefixTree {
    return NewPrefixTree()
}


type TrieNode struct {
	Children map[rune]*TrieNode
	EOW      bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{Children: make(map[rune]*TrieNode)}
}

type PrefixTree struct {
	root *TrieNode
}

func NewPrefixTree() PrefixTree {
	return PrefixTree{root: NewTrieNode()}
}

func (p *PrefixTree) Insert(word string) {
	curr := p.root
	for _, w := range word {
		if node, ok := curr.Children[w]; ok {
			curr = node
			continue
		}
		curr.Children[w] = NewTrieNode()
		curr = curr.Children[w]
	}
	curr.EOW = true
}

func (p *PrefixTree) Search(word string) bool {
	curr := p.root

	for _, w := range word {
		node, ok := curr.Children[w]
		if !ok {
			return false
		}
		curr = node
	}

	return curr.EOW
}

func (p *PrefixTree) StartsWith(prefix string) bool {
	curr := p.root

	for _, w := range prefix {
		node, ok := curr.Children[w]
		if !ok {
			return false
		}
		curr = node
	}

	return true
}
