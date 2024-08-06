package shared

import (
	"math/rand"
	"time"
)

type Board [3][3]string

type Game struct {
	Game             Board
	StartedTimestamp time.Time
	EndedTimestamp   *time.Time
	PlayerX          Player
	PlayerY          Player
	Players          []Player
	Turn             string
}

type Player struct {
	Id     string
	GameId string
}

var Games = make(map[string]*Game)
var Players = make(map[string]*Player)

func CheckWin(board Board) (string, bool) {
	for i := 0; i < 3; i++ {
		if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0], true
		}
		if board[0][i] != "" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i], true
		}
	}

	if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0], true
	}
	if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2], true
	}

	return "", false
}

func RenderBoard() {

}

func RandomizeTurn() string {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		return "X"
	}
	return "O"
}
