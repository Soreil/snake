package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

//directions the snake head is going
const (
	LEFT = iota
	UP
	RIGHT
	DOWN
)

type snake struct {
	length    int
	direction int
}

type board struct {
	width  int
	height int
	snakeX int
	snakeY int
	appleX int
	appleY int
}

//Create an apple at a random position
func (bPtr *board) newApple() {
	time := time.Now()
	rand.Seed(time.UnixNano())
	b := bPtr
	b.appleX = rand.Intn(b.width)
	b.appleY = rand.Intn(b.height)
	//TODO: change the next part to match all cells containing a snake part
	if b.appleX == b.snakeX && b.appleY == b.snakeY {
		b.newApple()
	} else {
		//draw apple
		termbox.SetCell(b.appleX, b.appleY, '*', termbox.ColorRed, termbox.ColorDefault)
	}
}

func moveSnake(sPtr *snake, bPtr *board) {
	snake := sPtr
	board := bPtr
	switch snake.direction {
	case LEFT:
		board.snakeX--
	case UP:
		board.snakeY--
	case RIGHT:
		board.snakeX++
	case DOWN:
		board.snakeY++
	}
	//Draw Snake starting position
	termbox.SetCell(board.snakeX, board.snakeY, '#', termbox.ColorDefault, termbox.ColorDefault)
}

func main() {
	//Create a window to draw in
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	//Close window at end of main
	defer termbox.Close()

	//Set up the snake
	snake := new(snake)
	snake.length = 4
	snake.direction = RIGHT

	//Set up the board
	board := new(board)
	board.width, board.height = termbox.Size()

	//Spawn Snake
	board.snakeX = board.width / 2
	board.snakeY = board.height / 2

	for i := 0; i < snake.length; i++ {
		moveSnake(snake, board)
	}

	//Spawn a Apple
	board.newApple()

	for board.snakeX <= board.width && board.snakeX >= 0 && board.snakeY <= board.height && board.snakeY >= 0 {
		break
		/* Game loop */
	}
	termbox.Flush()

	//Handle input, quit when Escape is pressed.
inputLoop:
	for {
		event := termbox.PollEvent()
		switch event.Type {
		case termbox.EventKey:
			if event.Key == termbox.KeyEsc {
				break inputLoop
			}
		}
	}
}
