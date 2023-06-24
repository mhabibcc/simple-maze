package main

import (
	"fmt"
)

func main() {
	var input string

	maze := [][]string{
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
	xAxis := 8
	yAxis := 15
	winPosX := 8
	winPosY := 0
inputLoop:
	for {
		ClearTerminal()
		showHeader()

		if xAxis == winPosX && yAxis == winPosY {
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("Yaay, Solved. Congratulations")
			break
		}

		for _, v := range maze {
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
			
		case "a", "A":
			fmt.Println("a key pressed")
			if maze[yAxis][xAxis-1] == " " {
				maze[yAxis][xAxis-1] = "@"
				maze[yAxis][xAxis] = " "
				xAxis--
			} else {
				fmt.Println("cannot move")
			}
		case "s", "S":
			fmt.Println("s key pressed")
			if maze[yAxis+1][xAxis] == " " {
				maze[yAxis+1][xAxis] = "@"
				maze[yAxis][xAxis] = " "
				yAxis++
			} else {
				fmt.Println("cannot move")
			}
		case "d", "D":
			fmt.Println("d key pressed")
			if maze[yAxis][xAxis+1] == " " {
				maze[yAxis][xAxis+1] = "@"
				maze[yAxis][xAxis] = " "
				xAxis++
			} else {
				fmt.Println("cannot move")
			}
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

/* 
 * will be modified later
 */

func moveUp(maze *[][]string) {
	
}

func moveDown(maze *[][]string)  {
	
}

func moveLeft(maze *[][]string)  {
	
}

func moveRight(maze *[][]string)  {
	
}