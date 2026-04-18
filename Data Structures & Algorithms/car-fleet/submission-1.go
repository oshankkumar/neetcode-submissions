
type Car struct {
	Position int
	Speed    int
}

type CarStack []Car

func (c *CarStack) Push(car Car) {
	*c = append(*c, car)
}

func (c *CarStack) Pop() Car {
	last := (*c)[len(*c)-1]
	*c = (*c)[:len(*c)-1]
	return last
}

func (c *CarStack) Top() Car {
	return (*c)[len(*c)-1]
}

func (c *CarStack) Len() int {
	return len(*c)
}

func carFleet(target int, position []int, speed []int) int {
	cars := make([]Car, len(position))

	for i := range position {
		cars[i] = Car{Position: position[i], Speed: speed[i]}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Position > cars[j].Position
	})

	var stk CarStack

	stk.Push(cars[0])

	for _, car := range cars[1:] {
		c1, c2 := stk.Top(), car
		t1, t2 := (float64(target)-float64(c1.Position))/float64(c1.Speed), (float64(target)-float64(c2.Position))/float64(c2.Speed)
		if t2 > t1 {
			stk.Push(c2)
		}
	}

	return stk.Len()
}
