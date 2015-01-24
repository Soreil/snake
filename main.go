package main

import (
	"errors"
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

type point struct {
	x int
	y int
}

type snake struct {
	length    int
	direction int
}

type board struct {
	width      int
	height     int
	snake      point
	apple      point
	snakeParts []point
}

//Create an apple at a random position
func (bPtr *board) newApple() {
	time := time.Now()
	rand.Seed(time.UnixNano())
	b := bPtr
	b.apple.x = rand.Intn(b.width)
	b.apple.y = rand.Intn(b.height)
	//check if the Apple will spawn inside of the snake, otherwise retry
	for _, v := range b.snakeParts {
		if v == b.apple {
			b.newApple()
		}
	}
	//draw apple
	termbox.SetCell(b.apple.x, b.apple.y, ' ', termbox.ColorDefault, termbox.ColorRed)
}

func moveSnake(sPtr *snake, bPtr *board) error {
	//error is nil
	var err error

	snake := sPtr
	board := bPtr
	//set direction
	switch snake.direction {
	case LEFT:
		board.snake.x--
	case UP:
		board.snake.y--
	case RIGHT:
		board.snake.x++
	case DOWN:
		board.snake.y++
	}

	//if the snake collides with itself return with an error
	for _, v := range board.snakeParts {
		if v == board.snake {
			err = errors.New("Snake collided with itself")
		}
	}

	board.snakeParts = append(board.snakeParts, board.snake)

	//do not remove the last part of the snake and increase the length
	if board.apple == board.snake {
		board.newApple()
		snake.length++
		//remove the last part of the snake
	} else if len(board.snakeParts) > 3 {
		termbox.SetCell(board.snakeParts[0].x, board.snakeParts[0].y, ' ', termbox.ColorDefault, termbox.ColorDefault)
		board.snakeParts = board.snakeParts[1:]
	}
	//Draw Snake starting position
	termbox.SetCell(board.snake.x, board.snake.y, ' ', termbox.ColorDefault, termbox.ColorWhite)
	return err
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
	board.width--
	board.height--

	//Spawn Snake
	board.snake.x = board.width / 2
	board.snake.y = board.height / 2

	for i := 0; i < snake.length; i++ {
		moveSnake(snake, board)
	}

	//Spawn an Apple
	board.newApple()

	//key presses
	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	//collect frame draw times
	frameDrawTimer := make(chan time.Time)
	go func() {
		for {
			frameDrawTimer <- <-time.After(time.Second / 15)
		}
	}()

	for board.snake.x <= board.width && board.snake.x >= 0 && board.snake.y <= board.height && board.snake.y >= 0 {
		//gameError := make(chan error)
		select {
		//If an arrow key has been pressed move the direction to the arrow key
		case event := <-event_queue:
			switch event.Type {
			case termbox.EventKey:
				switch event.Key {
				case termbox.KeyArrowDown:
					snake.direction = DOWN
				case termbox.KeyArrowLeft:
					snake.direction = LEFT
				case termbox.KeyArrowRight:
					snake.direction = RIGHT
				case termbox.KeyArrowUp:
					snake.direction = UP
				}
			}
			//After a frame time is done move and draw
		case <-frameDrawTimer:
			err := moveSnake(snake, board)
			termbox.Flush()
			if err != nil {
				board.snake.x = -1
				board.snake.y = -1
				//gameError <- err
			}
			//case err := <-gameError:
			//	fmt.Print(err)
		}
	}
	termbox.SetCell(0+board.width/2, board.height/2, 'G', termbox.ColorWhite, termbox.ColorRed)
	termbox.SetCell(1+board.width/2, board.height/2, 'A', termbox.ColorWhite, termbox.ColorRed)
	termbox.SetCell(2+board.width/2, board.height/2, 'M', termbox.ColorWhite, termbox.ColorRed)
	termbox.SetCell(3+board.width/2, board.height/2, 'E', termbox.ColorWhite, termbox.ColorRed)
	termbox.SetCell(4+board.width/2, board.height/2, ' ', termbox.ColorWhite, termbox.ColorRed)
	termbox.SetCell(5+board.width/2, board.height/2, 'O', termbox.ColorWhite, termbox.ColorRed)
	termbox.SetCell(6+board.width/2, board.height/2, 'V', termbox.ColorWhite, termbox.ColorRed)
	termbox.SetCell(7+board.width/2, board.height/2, 'E', termbox.ColorWhite, termbox.ColorRed)
	termbox.SetCell(8+board.width/2, board.height/2, 'R', termbox.ColorWhite, termbox.ColorRed)
	termbox.Flush()

	//rip game
inputLoop:
	for {
		event := <-event_queue
		switch event.Type {
		case termbox.EventKey:
			if event.Key == termbox.KeyEsc {
				break inputLoop
			}
		}
	}
}
