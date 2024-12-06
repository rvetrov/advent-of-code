package grid

import "fmt"

type Position struct {
	Row, Col int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.Row, p.Col)
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

func (d Direction) String() string {
	return fmt.Sprintf("<%d,%d>", d.DR, d.DC)
}

func (d Direction) Reversed() Direction {
	return Direction{-d.DR, -d.DC}
}

func (d Direction) Add(delta Direction) Direction {
	return Direction{d.DR + delta.DR, d.DC + delta.DC}
}

func (d Direction) Multiplied(x int) Direction {
	return Direction{d.DR * x, d.DC * x}
}

func (d Direction) AsPosition() Position {
	return Position{Row: d.DR, Col: d.DC}
}

func (d Direction) TurnCW() Direction {
	return Direction{
		DR: d.DC,
		DC: -d.DR,
	}
}

var (
	Up    = Direction{-1, 0}
	Down  = Direction{1, 0}
	Right = Direction{0, 1}
	Left  = Direction{0, -1}

	UpRight   = Up.Add(Right)
	UpLeft    = Up.Add(Left)
	DownRight = Down.Add(Right)
	DownLeft  = Down.Add(Left)

	FourSides = []Direction{Up, Right, Down, Left}

	EightSides = []Direction{Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft}
)
