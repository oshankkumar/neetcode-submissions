type MinStack struct {
	vals   []int
	minVal int
}

func Constructor() MinStack {
	return MinStack{minVal: math.MaxInt}
}

func (this *MinStack) Push(val int) {
	this.minVal = min(this.minVal, val)
	this.vals = append(this.vals, val)
}

func (this *MinStack) Pop() {
	if len(this.vals) == 0 {
		return
	}
	top := this.vals[len(this.vals)-1]
	this.vals = this.vals[:len(this.vals)-1]
	if top == this.minVal {
		this.minVal = math.MaxInt
		for _, v := range this.vals {
			this.minVal = min(this.minVal, v)
		}
	}
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
	return this.minVal
}
