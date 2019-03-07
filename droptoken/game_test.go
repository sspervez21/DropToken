package droptoken

import "testing"

func TestInitialize(t *testing.T) {
	player1 := "Player1"
	player2 := "Player2"
	columns := 4
	rows := 4

	game := createGame([2]string{player1, player2}, columns, rows)

	if game.nextMove != 0 {
		t.Fatalf("nextMove number should be 0, not: %d\n", game.nextMove)
	}

	if game.state != inProgress {
		t.Fatalf("the game should be in progress, not: %d\n", game.state)
	}

	if len(game.players) != 2 {
		t.Fatalf("Number of players incorrectly initialized.\n")
	}

	if game.players[0] != player1 {
		t.Fatalf("Player1 incorrectly initialized. %s\n", game.players[0])
	}

	if game.players[1] != player2 {
		t.Fatalf("Player2 incorrectly initialized. %s\n", game.players[1])
	}

	if game.curPlayer != player1 {
		t.Fatalf("It should be player1's turn, not %s\n", game.curPlayer)
	}

	if game.winner != "" {
		t.Fatalf("There should be no winner. %s\n", game.winner)
	}

	if game.columns != columns {
		t.Fatalf("Number of columns is incorrect. Found %d expected %d.", game.columns, columns)
	}

	if game.rows != rows {
		t.Fatalf("Number of rows is incorrect. Found %d expected %d.", game.rows, rows)
	}

	if len(game.moves) != 0 {
		t.Fatalf("There should be no moves made. %d\n", len(game.moves))
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			if game.board[row][col] != 0 {
				t.Fatalf("board not initialized correctly. board[%d][%d] == %d", row, col, game.board[row][col])
			}
		}
	}
}

func TestMakeMove(t *testing.T) {
	player1 := "Player1"
	player2 := "Player2"
	columns := 4
	rows := 4

	game := createGame([2]string{player1, player2}, columns, rows)

	game.makeMove(player1, 2)

	if len(game.moves) != 1 {
		t.Fatalf("move not included in list of moves. %d\n", len(game.moves))
	}

	if game.moves[0].moveNumber != game.nextMove-1 {
		t.Fatalf("moveNumber not updated correctly. %d:%d", game.moves[0].moveNumber, game.nextMove)
	}

	if game.moves[0].column != 2 {
		t.Fatalf("move.column incorrect. %d\n", game.moves[0].column)
	}

	if game.moves[0].player != player1 {
		t.Fatalf("move.player incorrect. %s\n", game.moves[0].player)
	}

	if game.curPlayer != player2 {
		t.Fatalf("player not updated after move.")
	}

	if game.state != inProgress {
		t.Fatalf("Game should be in progress. %d\n", game.state)
	}

	if game.winner != "" {
		t.Fatalf("Game should not have a winner. %s\n", game.winner)
	}
}

func TestMakeMoveGameOver1(t *testing.T) {
	player1 := "Player1"
	player2 := "Player2"
	columns := 4
	rows := 4

	game := createGame([2]string{player1, player2}, columns, rows)

	game.makeMove(player1, 2)
	game.makeMove(player2, 1)
	game.makeMove(player1, 2)
	game.makeMove(player2, 1)
	game.makeMove(player1, 2)
	game.makeMove(player2, 1)
	game.makeMove(player1, 2)

	if len(game.moves) != 7 {
		t.Fatalf("Incorrect number of moves. %d\n", len(game.moves))
	}

	if game.state != done {
		t.Fatalf("Game should be done. %d\n", game.state)
	}

	if game.winner != player1 {
		t.Fatalf("Game should be won by player1. %s\n", game.winner)
	}
}

func TestMakeMoveGameOver2(t *testing.T) {
	player1 := "Player1"
	player2 := "Player2"
	columns := 4
	rows := 4

	game := createGame([2]string{player1, player2}, columns, rows)

	game.makeMove(player1, 0)
	game.makeMove(player2, 0)
	game.makeMove(player1, 1)
	game.makeMove(player2, 1)
	game.makeMove(player1, 2)
	game.makeMove(player2, 2)
	game.makeMove(player1, 3)

	if len(game.moves) != 7 {
		t.Fatalf("Incorrect number of moves. %d\n", len(game.moves))
	}

	if game.state != done {
		t.Fatalf("Game should be done. %d\n", game.state)
	}

	if game.winner != player1 {
		t.Fatalf("Game should be won by player1. %s\n", game.winner)
	}
}

func TestMakeMoveGameOverDraw(t *testing.T) {
	player1 := "Player1"
	player2 := "Player2"
	columns := 4
	rows := 4

	game := createGame([2]string{player1, player2}, columns, rows)

	game.makeMove(player1, 0)
	game.makeMove(player2, 0)
	game.makeMove(player1, 1)
	game.makeMove(player2, 1)
	game.makeMove(player1, 2)
	game.makeMove(player2, 2)
	game.makeMove(player1, 3)

	if len(game.moves) != 7 {
		t.Fatalf("Incorrect number of moves. %d\n", len(game.moves))
	}

	if game.state != done {
		t.Fatalf("Game should be done. %d\n", game.state)
	}

	if game.winner != player1 {
		t.Fatalf("Game should be won by player1. %s\n", game.winner)
	}
}

func TestEndGameNegativeInit(t *testing.T) {

	game := &Game{}
	game.rows = 4
	game.columns = 4
	game.players = [2]string{"player1", "player2"}
	game.board = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	gameOver, winner := game.isGameOver()

	if gameOver {
		t.Fatalf("Game should not be over in this context\n")
	}

	if winner != "" {
		t.Fatalf("There should be no winner: %s\n", winner)
	}
}

func TestEndGameNegativeFilled(t *testing.T) {

	game := &Game{}
	game.rows = 4
	game.columns = 4
	game.players = [2]string{"player1", "player2"}
	game.board = [][]int{
		{2, 2, 2, 1},
		{0, 1, 1, 1},
		{0, 1, 0, 1},
		{0, 1, 0, 0},
	}

	gameOver, winner := game.isGameOver()

	if gameOver {
		t.Fatalf("Game should not be over in this context\n")
	}

	if winner != "" {
		t.Fatalf("There should be no winner: %s\n", winner)
	}
}

func TestEndGameColumnP1(t *testing.T) {

	game := &Game{}
	game.rows = 4
	game.columns = 4
	game.players = [2]string{"player1", "player2"}
	game.board = [][]int{
		{0, 1, 0, 0},
		{0, 1, 2, 0},
		{0, 1, 2, 0},
		{0, 1, 2, 0},
	}

	gameOver, winner := game.isGameOver()

	if !gameOver {
		t.Fatalf("Game should be over in this context\n")
	}

	if winner != "player1" {
		t.Fatalf("player1 should be the winner: %s\n", winner)
	}
}

func TestEndGameColumnP2(t *testing.T) {

	game := &Game{}
	game.rows = 4
	game.columns = 4
	game.players = [2]string{"player1", "player2"}
	game.board = [][]int{
		{0, 1, 2, 0},
		{0, 1, 2, 0},
		{0, 0, 2, 0},
		{0, 1, 2, 0},
	}

	gameOver, winner := game.isGameOver()

	if !gameOver {
		t.Fatalf("Game should be over in this context\n")
	}

	if winner != "player2" {
		t.Fatalf("player2 should be the winner: %s\n", winner)
	}
}

func TestEndGameRowP1(t *testing.T) {

	game := &Game{}
	game.rows = 4
	game.columns = 4
	game.players = [2]string{"player1", "player2"}
	game.board = [][]int{
		{0, 1, 2, 0},
		{1, 1, 1, 1},
		{0, 0, 2, 0},
		{0, 1, 2, 0},
	}

	gameOver, winner := game.isGameOver()

	if !gameOver {
		t.Fatalf("Game should be over in this context\n")
	}

	if winner != "player1" {
		t.Fatalf("player1 should be the winner: %s\n", winner)
	}
}

func TestEndGameRowP2(t *testing.T) {

	game := &Game{}
	game.rows = 4
	game.columns = 4
	game.players = [2]string{"player1", "player2"}
	game.board = [][]int{
		{0, 1, 2, 0},
		{1, 1, 2, 1},
		{0, 0, 2, 0},
		{2, 2, 2, 2},
	}

	gameOver, winner := game.isGameOver()

	if !gameOver {
		t.Fatalf("Game should be over in this context\n")
	}

	if winner != "player2" {
		t.Fatalf("player2 should be the winner: %s\n", winner)
	}
}

func TestEndGameDiag1(t *testing.T) {

	game := &Game{}
	game.rows = 4
	game.columns = 4
	game.players = [2]string{"player1", "player2"}
	game.board = [][]int{
		{1, 1, 2, 0},
		{1, 1, 2, 1},
		{0, 0, 1, 0},
		{2, 2, 2, 1},
	}

	gameOver, winner := game.isGameOver()

	if !gameOver {
		t.Fatalf("Game should be over in this context\n")
	}

	if winner != "player1" {
		t.Fatalf("player1 should be the winner: %s\n", winner)
	}
}

func TestEndGameDiag2(t *testing.T) {

	game := &Game{}
	game.rows = 4
	game.columns = 4
	game.players = [2]string{"player1", "player2"}
	game.board = [][]int{
		{1, 1, 2, 2},
		{1, 1, 2, 1},
		{0, 2, 0, 0},
		{2, 2, 2, 1},
	}

	gameOver, winner := game.isGameOver()

	if !gameOver {
		t.Fatalf("Game should be over in this context\n")
	}

	if winner != "player2" {
		t.Fatalf("player2 should be the winner: %s\n", winner)
	}
}

func TestEndGameDraw(t *testing.T) {

	game := &Game{}
	game.rows = 4
	game.columns = 4
	game.players = [2]string{"player1", "player2"}
	game.board = [][]int{
		{1, 1, 2, 2},
		{1, 1, 2, 1},
		{2, 1, 2, 1},
		{2, 2, 1, 1},
	}

	gameOver, winner := game.isGameOver()

	if !gameOver {
		t.Fatalf("Game should be over in this context\n")
	}

	if winner != "null" {
		t.Fatalf("There should be no winner: %s\n", winner)
	}
}
