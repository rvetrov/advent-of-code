package grid

type Position struct {
	Row, Col int
}

func (p Position) Add(d Direction) Position {
	return Position{p.Row + d.DR, p.Col + d.DC}
}

func (p Position) Subtract(other Position) Direction {
	return Direction{p.Row - other.Row, p.Col - other.Col}
}

type Direction struct {
	DR, DC int
}

func (v Direction) Reversed() Direction {
	return Direction{-v.DR, -v.DC}
}

var (
	Up    = Direction{-1, 0}
	Down  = Direction{1, 0}
	Right = Direction{0, 1}
	Left  = Direction{0, -1}
)
