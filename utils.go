package main

type Point struct {
	X float64
	Y float64
}

type Direction struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

type Velocity struct {
	X float64
	Y float64
}

func (v *Velocity) SetFromDirection(speed float64, up, down, left, right bool) {
	v.X, v.Y = 0, 0
	if left {
		v.X = -speed
	} else if right {
		v.X = speed
	}
	if up {
		v.Y = -speed
	} else if down {
		v.Y = speed
	}
}
