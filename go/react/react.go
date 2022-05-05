package react

type cell struct {
	value     int
	reactor   *reactor
	dep1      *Cell
	dep2      *Cell
	func1     func(int) int
	func2     func(int, int) int
	callbacks []*func(int)
}
type reactor struct {
	computeCells []*cell
}
type canceler struct {
	cell     *cell
	callback *func(int)
}

func (c *canceler) Cancel() {
	for index, callback := range c.cell.callbacks {
		if callback == c.callback {
			c.cell.callbacks = append(c.cell.callbacks[:index], c.cell.callbacks[index+1:]...)
			return
		}
	}
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) SetValue(value int) {
	c.value = value
	for _, computeCell := range c.reactor.computeCells {
		var newValue int
		if computeCell.func1 != nil {
			newValue = computeCell.func1((*(computeCell.dep1)).Value())
		} else {
			newValue = computeCell.func2((*(computeCell.dep1)).Value(), (*(computeCell.dep2)).Value())
		}
		if newValue != computeCell.value {
			computeCell.value = newValue
			for _, callback := range computeCell.callbacks {
				(*callback)(computeCell.value)
			}
		}
	}
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	c.callbacks = append(c.callbacks, &callback)
	return &canceler{cell: c, callback: &callback}
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(initial int) InputCell {
	cell := cell{value: initial, reactor: r}
	return &cell

}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	cell := cell{value: compute(dep.Value()), reactor: r, dep1: &dep, func1: compute}
	r.computeCells = append(r.computeCells, &cell)
	return &cell
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	cell := cell{value: compute(dep1.Value(), dep2.Value()), reactor: r, dep1: &dep1, dep2: &dep2, func2: compute}
	r.computeCells = append(r.computeCells, &cell)
	return &cell
}
