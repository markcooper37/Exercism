package clock

import "fmt"

type Clock struct {
    Hour int
    Minute int
}

func New(h, m int) Clock {
    for m >= 60 {
        m -= 60
        h += 1
    }
	for m < 0 {
        m += 60
        h -= 1
    }
	h = ((h % 24) + 24) % 24
    clock := Clock{Hour: h, Minute: m}
	return clock
}

func (c Clock) Add(m int) Clock {
	return New(c.Hour, c.Minute + m)
}

func (c Clock) Subtract(m int) Clock {
    return New(c.Hour, c.Minute - m)
}

func (c Clock) String() string {
    return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}
