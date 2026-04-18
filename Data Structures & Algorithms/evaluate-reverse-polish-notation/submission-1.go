


type StringStack []string

func (ss *StringStack) Push(val string) {
	*ss = append(*ss, val)
}

func (ss *StringStack) Top() string {
	return (*ss)[len(*ss)-1]
}

func (ss *StringStack) Pop() string {
	n, last := len(*ss), (*ss)[len(*ss)-1]
	*ss = (*ss)[:n-1]
	return last
}

func (ss *StringStack) Len() int {
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
	var evalStack StringStack
	for _, token := range tokens {
		op, ok := Operators[token]
		if !ok {
			evalStack.Push(token)
			continue
		}

		oper1, oper2 := mustInt(evalStack.Pop()), mustInt(evalStack.Pop())

		evalStack.Push(
			strconv.Itoa(
				op(oper2, oper1),
			),
		)
	}

	return mustInt(evalStack.Pop())
}
