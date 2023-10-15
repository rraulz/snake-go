package snake

const (
	RIGHT direction = 1 + iota
	LEFT
	UP
	DOWN
	NONE
)

type direction int

type Snake struct {
	body           []Coord
	length         int
	snakeDirection direction
}

func NewSnake() *Snake {
	return &Snake{snakeDirection: RIGHT, length: 1}
}

func (s *Snake) getHeadSnake() Coord {
	return s.body[len(s.body)-1]
}

func (s *Snake) placeSnakeHeadMiddleOfArea(a *Arena) {
	middleCoord := Coord{((a.getTopRightCorner().x - a.getTopLeftCorner().x) / 2) + a.leftMargin, ((a.getBottomLeftCorner().y - a.getTopLeftCorner().y) / 2) + a.topMargin}

	middleCoord2 := Coord{((a.getTopRightCorner().x - a.getTopLeftCorner().x) / 2) + a.leftMargin - 1, ((a.getBottomLeftCorner().y - a.getTopLeftCorner().y) / 2) + a.topMargin}

	s.length = 2
	s.body = append(s.body, middleCoord, middleCoord2)
}

func (s *Snake) changeSnakeDirection(d direction) {
	opposites := map[direction]direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}
	if o := opposites[d]; o != 0 && o != s.snakeDirection {
		s.snakeDirection = d
	}

}

func (s *Snake) moveSnake(a *Arena) {

	if canMakeTheMove, newCoord := isNextSnakeMoveAllowed(s, a); canMakeTheMove {

		if s.length > len(s.body) {
			s.body = append(s.body, newCoord)
		} else {
			s.body = append(s.body[1:], newCoord)
		}
	}
}

func isNextSnakeMoveAllowed(s *Snake, a *Arena) (bool, Coord) {

	newCoord := getNextMoveCoordSnakeHead(s)

	if newCoord.x == a.getTopLeftCorner().x || newCoord.x == a.getTopRightCorner().x || newCoord.y == a.getTopRightCorner().y || newCoord.y == a.getBottomRightCorner().y {
		return false, Coord{0, 0}
	}

	return true, newCoord
}

func getNextMoveCoordSnakeHead(s *Snake) Coord {
	newX := s.getHeadSnake().x
	newY := s.getHeadSnake().y
	switch s.snakeDirection {
	case RIGHT:
		newX++
	case LEFT:
		newX--
	case UP:
		newY--
	case DOWN:
		newY++
	}
	return Coord{newX, newY}
}
