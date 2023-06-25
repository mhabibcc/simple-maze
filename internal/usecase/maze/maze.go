package maze

import "fmt"

type Maze struct {
	Maze [][]string
	X    int
	Y    int
	WinX int
	WinY int
}

func NewMaze(initMaze [][]string, xAxis, yAxis, xWinPos, yWinPos int) *Maze {
	return &Maze{
		Maze: initMaze,
		X: xAxis,
		Y: yAxis,
		WinX: xWinPos,
		WinY: yWinPos,
	}
}

func (m *Maze) MoveUp() {
	if m.Maze[m.Y-1][m.X] == " " {
		m.Maze[m.Y-1][m.X] = "@"
		m.Maze[m.Y][m.X] = " "
		m.Y--
	} else {
		fmt.Println("cannot move")
	}
}

func (m *Maze) MoveDown() {
	if m.Maze[m.Y+1][m.X] == " " {
		m.Maze[m.Y+1][m.X] = "@"
		m.Maze[m.Y][m.X] = " "
		m.Y++
	} else {
		fmt.Println("cannot move")
	}
}

func (m *Maze) MoveLeft() {
	if m.Maze[m.Y][m.X-1] == " " {
		m.Maze[m.Y][m.X-1] = "@"
		m.Maze[m.Y][m.X] = " "
		m.X--
	} else {
		fmt.Println("cannot move")
	}
}

func (m *Maze) MoveRight() {
	if m.Maze[m.Y][m.X+1] == " " {
		m.Maze[m.Y][m.X+1] = "@"
		m.Maze[m.Y][m.X] = " "
		m.X++
	} else {
		fmt.Println("cannot move")
	}
}

func (m *Maze) IsWin() bool {
	return m.X == m.WinX && m.Y == m.WinY
}
