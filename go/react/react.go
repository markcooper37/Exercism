package react

type cell struct {
	value        int
	reactor      *reactor
	computeValue func() int
	callbacks    map[int]func(int)
}
type reactor struct {
	computeCells []*cell
}
type canceler struct {
	cell *cell
	id   int
}

func (c *canceler) Cancel() {
	delete(c.cell.callbacks, c.id)
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) SetValue(value int) {
	c.value = value
	for _, computeCell := range c.reactor.computeCells {
		if newValue := computeCell.computeValue(); newValue != computeCell.value {
			computeCell.value = newValue
			for _, callback := range computeCell.callbacks {
				callback(computeCell.value)
			}
		}
	}
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	for i := 0; i <= len(c.callbacks); i++ {
		if _, ok := c.callbacks[i]; !ok {
			c.callbacks[i] = callback
			return &canceler{cell: c, id: i}
		}
	}
	return nil
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(initial int) InputCell {
	return &cell{value: initial, reactor: r}
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	computeValue := func() int { return compute(dep.Value()) }
	cell := cell{value: computeValue(), computeValue: computeValue, reactor: r, callbacks: map[int]func(int){}}
	r.computeCells = append(r.computeCells, &cell)
	return &cell
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	computeValue := func() int { return compute(dep1.Value(), dep2.Value()) }
	cell := cell{value: computeValue(), computeValue: computeValue, reactor: r, callbacks: map[int]func(int){}}
	r.computeCells = append(r.computeCells, &cell)
	return &cell
}
