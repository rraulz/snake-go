package snake

import (
	"time"

	"github.com/nsf/termbox-go"
)

type Game struct {
	score int
	snake *Snake
	arena *Arena
	timer *Timer
}

var (
	keyboardEventsChan = make(chan keyboardEvent)
)

func initialScore() int {
	return 0
}

func NewGame() *Game {
	return &Game{score: initialScore(), snake: NewSnake(), arena: NewArena(15, 20, 2, 2), timer: NewTimer()}
}

func (g *Game) moveInterval() time.Duration {
	ms := 100 - (g.score / 10)
	return time.Duration(ms) * time.Millisecond
}

func snakeMovement(g *Game) {
	for {
		g.snake.moveSnake(g.arena)
		time.Sleep(g.moveInterval())
	}
}

func (g *Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	g.snake.placeSnakeHeadMiddleOfArea(g.arena)

	go listenToKeyboard(keyboardEventsChan)
	go snakeMovement(g)

mainloop:
	for {
		select {
		case p := <-keyboardEventsChan:
			switch p.eventType {
			case MOVE:
				g.snake.changeSnakeDirection(p.action)
			case TEST:
				g.snake.length++
			case END:
				break mainloop
			default:

			}
		default:
			g.render()
		}

	}
}
