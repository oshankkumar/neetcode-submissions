
type Parenthesis byte

func (p Parenthesis) Closing() Parenthesis {
	return map[Parenthesis]Parenthesis{
		'{': '}',
		'(': ')',
		'[': ']',
	}[p]
}

type PStack []Parenthesis

func (b *PStack) Push(d Parenthesis) {
	*b = append(*b, d)
}

func (b *PStack) Pop() Parenthesis {
	old, n := *b, len(*b)
	val := old[n-1]
	*b = old[:n-1]
	return val
}

func (b *PStack) Len() int {
	return len(*b)
}

func (b *PStack) Top() Parenthesis {
	return (*b)[len(*b)-1]
}

func isValid(s string) bool {
	var stack PStack

	expr := []Parenthesis(s)
	if len(expr)%2 != 0 {
		return false
	}

	for _, b := range expr {
		if stack.Len() == 0 {
			stack.Push(b)
			continue
		}
		if b == stack.Top().Closing() {
			stack.Pop()
			continue
		}

		stack.Push(b)
	}

	return stack.Len() == 0
}
