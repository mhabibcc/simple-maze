package maze

import (
	"reflect"
	"testing"
)

func TestNewMaze(t *testing.T) {
	initMaze := [][]string{
		{"#", "#", "#"},
		{"#", "#", "#"},
		{"#", "#", "#"},
	}
	type args struct {
		initMaze [][]string
		xAxis    int
		yAxis    int
		xWinPos  int
		yWinPos  int
	}
	tests := []struct {
		name string
		args args
		want *Maze
	}{
		{
			name: "initial maze",
			args: args{
				initMaze: initMaze,
				xAxis:    1,
				yAxis:    2,
				xWinPos:  1,
				yWinPos:  0,
			},
			want: &Maze{
				Maze: initMaze,
				X:    1,
				Y:    2,
				WinX: 1,
				WinY: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMaze(tt.args.initMaze, tt.args.xAxis, tt.args.yAxis, tt.args.xWinPos, tt.args.yWinPos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMaze() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaze_MoveUp(t *testing.T) {
	initMaze := [][]string{
		{"#", "#", "#"},
		{"#", " ", "#"},
		{"#", "@", "#"},
	}
	upMaze := [][]string{
		{"#", "#", "#"},
		{"#", "@", "#"},
		{"#", " ", "#"},
	}
	type fields struct {
		Maze [][]string
		X    int
		Y    int
		WinX int
		WinY int
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name: "valid move",
			fields: fields{
				Maze: initMaze,
				X:    1,
				Y:    2,
				WinX: 1,
				WinY: 1,
			},
			want: fields{
				Maze: upMaze,
				X:    1,
				Y:    1,
				WinX: 1,
				WinY: 1,
			},
		},
		{
			name: "invalid move",
			fields: fields{
				Maze: upMaze,
				X:    1,
				Y:    1,
				WinX: 1,
				WinY: 1,
			},
			want: fields{
				Maze: upMaze,
				X:    1,
				Y:    1,
				WinX: 1,
				WinY: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Maze{
				Maze: tt.fields.Maze,
				X:    tt.fields.X,
				Y:    tt.fields.Y,
				WinX: tt.fields.WinX,
				WinY: tt.fields.WinY,
			}
			m.MoveUp()
			if !reflect.DeepEqual(m.Maze, tt.want.Maze) || !reflect.DeepEqual(m.Y, tt.want.Y) {
				t.Errorf("maze = %v, want %v", *m, tt.want)
			}
		})
	}
}

func TestMaze_MoveDown(t *testing.T) {
	initMaze := [][]string{
		{"#", "#", "#"},
		{"#", "@", "#"},
		{"#", " ", "#"},
		{"#", "#", "#"},
	}
	downMaze := [][]string{
		{"#", "#", "#"},
		{"#", " ", "#"},
		{"#", "@", "#"},
		{"#", "#", "#"},
	}
	type fields struct {
		Maze [][]string
		X    int
		Y    int
		WinX int
		WinY int
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name: "valid move",
			fields: fields{
				Maze: initMaze,
				X:    1,
				Y:    1,
				WinX: 1,
				WinY: 1,
			},
			want: fields{
				Maze: downMaze,
				X:    1,
				Y:    2,
				WinX: 1,
				WinY: 1,
			},
		},
		{
			name: "invalid move",
			fields: fields{
				Maze: downMaze,
				X:    1,
				Y:    2,
				WinX: 1,
				WinY: 1,
			},
			want: fields{
				Maze: downMaze,
				X:    1,
				Y:    2,
				WinX: 1,
				WinY: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Maze{
				Maze: tt.fields.Maze,
				X:    tt.fields.X,
				Y:    tt.fields.Y,
				WinX: tt.fields.WinX,
				WinY: tt.fields.WinY,
			}
			m.MoveDown()
			if !reflect.DeepEqual(m.Maze, tt.want.Maze) || !reflect.DeepEqual(m.Y, tt.want.Y) {
				t.Errorf("maze = %v, want %v", *m, tt.want)
			}
		})
	}
}

func TestMaze_MoveLeft(t *testing.T) {
	initMaze := [][]string{
		{"#", "#", "#", "#"},
		{"#", " ", "@", "#"},
		{"#", "#", "#", "#"},
	}
	leftMaze := [][]string{
		{"#", "#", "#", "#"},
		{"#", "@", " ", "#"},
		{"#", "#", "#", "#"},
	}
	type fields struct {
		Maze [][]string
		X    int
		Y    int
		WinX int
		WinY int
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name: "valid move",
			fields: fields{
				Maze: initMaze,
				X:    2,
				Y:    1,
			},
			want: fields{
				Maze: leftMaze,
				X:    1,
				Y:    1,
			},
		},
		{
			name: "invalid move",
			fields: fields{
				Maze: leftMaze,
				X:    1,
				Y:    1,
			},
			want: fields{
				Maze: leftMaze,
				X:    1,
				Y:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Maze{
				Maze: tt.fields.Maze,
				X:    tt.fields.X,
				Y:    tt.fields.Y,
				WinX: tt.fields.WinX,
				WinY: tt.fields.WinY,
			}
			m.MoveLeft()
			if !reflect.DeepEqual(m.Maze, tt.want.Maze) || !reflect.DeepEqual(m.Y, tt.want.Y) {
				t.Errorf("maze = %v, want %v", *m, tt.want)
			}
		})
	}
}

func TestMaze_MoveRight(t *testing.T) {
	initMaze := [][]string{
		{"#", "#", "#", "#"},
		{"#", "@", " ", "#"},
		{"#", "#", "#", "#"},
	}
	rightMaze := [][]string{
		{"#", "#", "#", "#"},
		{"#", " ", "@", "#"},
		{"#", "#", "#", "#"},
	}
	type fields struct {
		Maze [][]string
		X    int
		Y    int
		WinX int
		WinY int
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name: "valid move",
			fields: fields{
				Maze: initMaze,
				X:    1,
				Y:    1,
			},
			want: fields{
				Maze: rightMaze,
				X:    2,
				Y:    1,
			},
		},
		{
			name: "invalid move",
			fields: fields{
				Maze: rightMaze,
				X:    2,
				Y:    1,
			},
			want: fields{
				Maze: rightMaze,
				X:    2,
				Y:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Maze{
				Maze: tt.fields.Maze,
				X:    tt.fields.X,
				Y:    tt.fields.Y,
				WinX: tt.fields.WinX,
				WinY: tt.fields.WinY,
			}
			m.MoveRight()
		})
	}
}

func TestMaze_IsWin(t *testing.T) {
	type fields struct {
		Maze [][]string
		X    int
		Y    int
		WinX int
		WinY int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "is win",
			fields: fields{
				X: 1,
				Y: 2,
				WinX: 1,
				WinY: 2,
			},
			want: true,
		},
		{
			name: "not win",
			fields: fields{
				X: 1,
				Y: 2,
				WinX: 3,
				WinY: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Maze{
				Maze: tt.fields.Maze,
				X:    tt.fields.X,
				Y:    tt.fields.Y,
				WinX: tt.fields.WinX,
				WinY: tt.fields.WinY,
			}
			if got := m.IsWin(); got != tt.want {
				t.Errorf("Maze.IsWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
