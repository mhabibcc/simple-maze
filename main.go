package main

import (
	"fmt"
)

type Maze struct {
	Maze [][]string
	X    int
	Y    int
	WinX int
	WinY int
}

func main() {
	var input string

	initMaze := [][]string{
		{"#", "#", "#", "#", "#", "#", "#", "#", " ", "#"},
		{"#", " ", " ", " ", " ", " ", " ", "#", " ", "#"},
		{"#", " ", "#", "#", "#", " ", "#", " ", " ", "#"},
		{"#", " ", "#", "#", " ", " ", "#", " ", "#", "#"},
		{"#", " ", " ", "#", " ", "#", "#", " ", "#", "#"},
		{"#", " ", " ", "#", " ", " ", " ", " ", "#", "#"},
		{"#", "#", " ", "#", "#", "#", "#", "#", "#", "#"},
		{"#", "#", " ", " ", " ", " ", " ", " ", "#", "#"},
		{"#", "#", "#", "#", "#", "#", "#", " ", "#", "#"},
		{"#", " ", " ", " ", " ", " ", "#", " ", " ", "#"},
		{"#", "#", "#", " ", "#", " ", "#", "#", " ", "#"},
		{"#", " ", "#", " ", "#", " ", " ", " ", " ", "#"},
		{"#", " ", " ", " ", "#", "#", "#", "#", " ", "#"},
		{"#", "#", "#", " ", " ", " ", " ", "#", " ", "#"},
		{"#", " ", " ", " ", "#", "#", " ", "#", " ", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#", "@", "#"},
	}

	maze := Maze{
		Maze: initMaze,
		X:    8,
		Y:    15,
		WinX: 8,
		WinY: 0,
	}

inputLoop:
	for {
		ClearTerminal()
		showHeader()

		if maze.X == maze.WinX && maze.Y == maze.WinY {
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("Yaay, Solved. Congratulations")
			break
		}

		for _, v := range maze.Maze {
			for _, s := range v {
				fmt.Print(s)
			}
			fmt.Println()
		}

		fmt.Scanln(&input)

		switch input {
		case "q", "Q":
			fmt.Println("               Exit Program. Thank You")
			break inputLoop
		case "w", "W":
			fmt.Println("w key pressed")
			maze.moveUp()
		case "a", "A":
			fmt.Println("a key pressed")
			maze.moveLeft()
		case "s", "S":
			fmt.Println("s key pressed")
			maze.moveDown()
		case "d", "D":
			fmt.Println("d key pressed")
			maze.moveRight()
		default:
			fmt.Println("invalid Key Pressed")
		}

	}
}

func showHeader() {
	fmt.Println("@----------------------------### SIMPLE MAZE ###-----------------------------@")
	fmt.Println("                      ---- Welcome to Simple Maze ----")
	fmt.Println()
	fmt.Println("               Type W/A/S/D and then press ENTER to move the `@`.")
	fmt.Println("               Type Q and press ENTER to exit")
}

func (m *Maze) moveUp() {
	if m.Maze[m.Y-1][m.X] == " " {
		m.Maze[m.Y-1][m.X] = "@"
		m.Maze[m.Y][m.X] = " "
		m.Y--
	} else {
		fmt.Println("cannot move")
	}
}

func (m *Maze) moveDown() {
	if m.Maze[m.Y+1][m.X] == " " {
		m.Maze[m.Y+1][m.X] = "@"
		m.Maze[m.Y][m.X] = " "
		m.Y++
	} else {
		fmt.Println("cannot move")
	}
}

func (m *Maze) moveLeft() {
	if m.Maze[m.Y][m.X-1] == " " {
		m.Maze[m.Y][m.X-1] = "@"
		m.Maze[m.Y][m.X] = " "
		m.X--
	} else {
		fmt.Println("cannot move")
	}
}

func (m *Maze) moveRight() {
	if m.Maze[m.Y][m.X+1] == " " {
		m.Maze[m.Y][m.X+1] = "@"
		m.Maze[m.Y][m.X] = " "
		m.X++
	} else {
		fmt.Println("cannot move")
	}
}
