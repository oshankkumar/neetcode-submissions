type IntStack []int

func (ss *IntStack) Push(val int) {
	*ss = append(*ss, val)
}

func (ss *IntStack) Top() int {
	return (*ss)[len(*ss)-1]
}

func (ss *IntStack) Pop() int {
	n, last := len(*ss), (*ss)[len(*ss)-1]
	*ss = (*ss)[:n-1]
	return last
}

func (ss *IntStack) Len() int {
	return len(*ss)
}

type Operator func(a, b int) int

var Operators = map[string]Operator{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func mustInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func evalRPN(tokens []string) int {
	var evalStack IntStack
	for _, token := range tokens {
		op, ok := Operators[token]
		if !ok {
			evalStack.Push(mustInt(token))
			continue
		}

		oper1, oper2 := evalStack.Pop(), evalStack.Pop()

		evalStack.Push(op(oper2, oper1))
	}

	return evalStack.Pop()
}
