package snake

import "github.com/nsf/termbox-go"

type keyboardEventType int

const (
	MOVE keyboardEventType = 1 + iota
	RETRY
	END
	TEST
	NOT_CONFIGURED
)

type keyboardEvent struct {
	eventType keyboardEventType
	action    direction
}

func listenToKeyboard(evChan chan keyboardEvent) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch keyEvent := termbox.PollEvent(); keyEvent.Type {
		case termbox.EventKey:
			evChan <- keyToEventParser(keyEvent.Ch)
		case termbox.EventError:
			panic(keyEvent.Err)
		}
	}
}

func keyToEventParser(keyPressed rune) keyboardEvent {

	switch keyPressed {
	case 'w':
		return keyboardEvent{eventType: MOVE, action: UP}
	case 's':
		return keyboardEvent{eventType: MOVE, action: DOWN}
	case 'a':
		return keyboardEvent{eventType: MOVE, action: LEFT}
	case 'd':
		return keyboardEvent{eventType: MOVE, action: RIGHT}
	case 'r':
		return keyboardEvent{eventType: RETRY, action: NONE}
	case 'x':
		return keyboardEvent{eventType: END, action: NONE}
	case 'f':
		return keyboardEvent{eventType: TEST, action: NONE}
	default:
		return keyboardEvent{eventType: NOT_CONFIGURED, action: NONE}
	}

}
