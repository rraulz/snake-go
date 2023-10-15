package snake

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

func renderSnake(s *Snake) {

	for i, bodyCoord := range s.body {
		if i+1 == len(s.body) {
			termbox.SetCell(bodyCoord.x, bodyCoord.y, ' ', termbox.ColorGreen, termbox.ColorGreen)
		} else {
			termbox.SetCell(bodyCoord.x, bodyCoord.y, ' ', termbox.ColorCyan, termbox.ColorCyan)
		}
	}
}

func renderArena(a *Arena) {

	//lineas inferior y posterior
	fill(a.getTopLeftCorner(), a.getTopRightCorner(), '─')
	fill(a.getBottomLeftCorner(), a.getBottomRightCorner(), '─')

	//lineas izquierdas y derechas
	fill(a.getTopLeftCorner(), a.getBottomLeftCorner(), '│')
	fill(a.getTopRightCorner(), a.getBottomRightCorner(), '│')

	//Esquinas
	termbox.SetCell(0+a.leftMargin, 0+a.topMargin, '┌', termbox.ColorBlue, termbox.ColorDefault)
	termbox.SetCell(0+a.leftMargin, a.height+a.topMargin, '└', termbox.ColorBlue, termbox.ColorDefault)
	termbox.SetCell(a.width+a.leftMargin, 0+a.topMargin, '┐', termbox.ColorBlue, termbox.ColorDefault)
	termbox.SetCell(a.width+a.leftMargin, a.height+a.topMargin, '┘', termbox.ColorBlue, termbox.ColorDefault)
}

func fill(pointA, pointB Coord, char rune) {

	if pointA.x > pointB.x || pointA.y > pointB.y {
		return
	}

	for i := pointA.x; i <= pointB.x; i++ {
		for u := pointA.y; u <= pointB.y; u++ {
			termbox.SetCell(i, u, char, termbox.ColorBlue, termbox.ColorDefault)
		}
	}
}

func renderScore(score, leftMargin int) {

	sScore := fmt.Sprintf("%d", score)
	runeScoreArray := []rune(sScore)

	for i, s := range runeScoreArray {
		termbox.SetCell(i+leftMargin, 0, s, termbox.ColorRed, termbox.ColorDefault)
	}
}

func renderTimer(t *Timer, leftMargin int) {

	seconds := int(time.Since(t.startTime).Seconds())
	sSeconds := fmt.Sprintf("%d", seconds)
	runeTimeArray := []rune(sSeconds)

	for i, s := range runeTimeArray {
		termbox.SetCell(i+leftMargin-len(runeTimeArray), 0, s, termbox.ColorRed, termbox.ColorDefault)
	}
}

func (g *Game) render() {

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	renderArena(g.arena)
	renderSnake(g.snake)
	renderTimer(g.timer, g.arena.width+g.arena.leftMargin)
	renderScore(g.score, g.arena.leftMargin)
	termbox.Flush()
}
