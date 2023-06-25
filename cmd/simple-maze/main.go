package main

import (
	"fmt"
	"os"

	"github.com/joking/simple-maze/internal/usecase/maze"
	oscommand "github.com/joking/simple-maze/pkg/os-command"
)

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

	maze := maze.NewMaze(initMaze, 8, 15, 8, 0)

	fmt.Println(*maze)
	os.Exit(0)

inputLoop:
	for {
		oscommand.ClearTerminal()
		showHeader()

		if maze.IsWin() {
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
			maze.MoveUp()
		case "a", "A":
			maze.MoveLeft()
		case "s", "S":
			maze.MoveDown()
		case "d", "D":
			maze.MoveRight()
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
