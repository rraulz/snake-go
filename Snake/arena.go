package snake

type Arena struct {
	height, width         int
	topMargin, leftMargin int
}

type Coord struct {
	x, y int
}

func NewArena(height, width, topMargin, leftMargin int) *Arena {
	return &Arena{height, width, topMargin, leftMargin}
}

func (a *Arena) getTopLeftCorner() Coord {

	return Coord{a.leftMargin, a.topMargin}
}

func (a *Arena) getTopRightCorner() Coord {

	return Coord{a.leftMargin + a.width, a.topMargin}
}

func (a *Arena) getBottomLeftCorner() Coord {

	return Coord{a.leftMargin, a.topMargin + a.height}
}

func (a *Arena) getBottomRightCorner() Coord {

	return Coord{a.leftMargin + a.width, a.topMargin + a.height}
}
