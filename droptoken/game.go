package droptoken

import "strconv"

// GameState enum
type GameState int

const (
	inProgress GameState = 0
	done       GameState = 1
)

// Game represents the state of a single Drop token game
type Game struct {
	gameID    string
	nextMove  int // ID of next move in game
	state     GameState
	players   [2]string
	curPlayer string // whose move is it
	winner    string
	columns   int
	rows      int
	moves     map[int]Move
	board     [][]int
}

// TODO: This can be turned into a persistent store
var allGamesMap = make(map[string]*Game)
var nextGameID = 0

func createGame(_players [2]string, _columns int, _rows int) *Game {
	// assert players.Length == 2

	gameID := "Game" + strconv.Itoa(nextGameID)

	game := Game{
		gameID:    gameID,
		nextMove:  0,
		state:     inProgress,
		players:   [2]string{_players[0], _players[1]},
		curPlayer: _players[0], // always start with player1
		winner:    "",
		columns:   _columns,
		rows:      _rows,
		moves:     make(map[int]Move),
	}

	game.board = make([][]int, _rows)
	for col := 0; col < _columns; col++ {
		game.board[col] = make([]int, _columns)
	}

	allGamesMap[gameID] = &game
	nextGameID++

	return &game
}

func getAllGames() []string {
	var games = []string{}
	for _, value := range allGamesMap {
		if value.state == inProgress {
			games = append(games, value.gameID)
		}
	}

	return games
}

func getGame(gameID string) (*Game, bool) {
	game, ok := allGamesMap[gameID]
	return game, ok
}

func (game *Game) isPlayerInGame(playerID string) bool {
	for _, player := range game.players {
		if player == playerID {
			return true
		}
	}

	return false
}

// assume game is in progress, player is part of the game and it is player's turn
func (game *Game) isLegalMove(playerID string, column int) bool {
	for row := 0; row < game.rows; row++ {
		if game.board[row][column] == 0 {
			return true
		}
	}

	return false
}

// assume move is legal
func (game *Game) makeMove(playerID string, column int) int {
	// TODO: double check move is legal

	m := Move{moveNumber: game.nextMove, mType: move, player: playerID, column: column}
	game.nextMove++
	game.curPlayer = game.getNextPlayer(playerID)
	game.moves[m.moveNumber] = m

	found := false
	for row := 0; row < game.rows; row++ {
		if game.board[row][column] == 0 {
			game.board[row][column] = game.getPlayerIndex(playerID)
			found = true
			break
		}
	}

	if !found {
		// indicate error state
	}

	// determine if game is over/winner
	gameOver, winner := game.isGameOver()

	if gameOver {
		game.state = done
		game.winner = winner
	}

	return m.moveNumber
}

// assume player is in game
func (game *Game) removePlayer(playerID string) {
	m := Move{moveNumber: game.nextMove, mType: quit, player: playerID}
	game.nextMove++
	game.moves[m.moveNumber] = m

	game.state = done
	game.winner = game.getNextPlayer(playerID)
	game.curPlayer = ""
}

// returns whether game is over, if it is optionally returns the winner
func (game *Game) isGameOver() (bool, string) {

	player1Count, player2Count := 0, 0
	player1TotalCount, player2TotalCount := 0, 0

	// check columns
	for col := 0; col < game.columns; col++ {
		player1Count, player2Count = game.getPlayerCount(0, col, 1, 0, game.rows, col)
		if player1Count == game.rows {
			return true, game.players[0]
		}
		if player2Count == game.rows {
			return true, game.players[1]
		}
		player1TotalCount += player1Count
		player2TotalCount += player2Count
	}

	// check rows
	for row := 0; row < game.rows; row++ {
		player1Count, player2Count = game.getPlayerCount(row, 0, 0, 1, row, game.columns)
		if player1Count == game.columns {
			return true, game.players[0]
		}
		if player2Count == game.columns {
			return true, game.players[1]
		}
	}

	// check diagonals
	player1Count, player2Count = game.getPlayerCount(0, 0, 1, 1, game.rows, game.columns)
	if player1Count == game.columns {
		return true, game.players[0]
	}
	if player2Count == game.columns {
		return true, game.players[1]
	}

	player1Count, player2Count = game.getPlayerCount(game.rows-1, 0, -1, 1, -1, game.columns)
	if player1Count == game.columns {
		return true, game.players[0]
	}
	if player2Count == game.columns {
		return true, game.players[1]
	}

	// None of the players won outright, check for a draw
	if player1TotalCount+player2TotalCount == game.rows*game.columns {
		return true, "null"
	}

	return false, ""
}

func (game *Game) getPlayerCount(startRow int, startCol int, stepRow int, stepCol int, endRow int, endCol int) (int, int) {
	player1Count, player2Count := 0, 0

	for startCol != endCol || startRow != endRow {
		if game.board[startRow][startCol] == game.getPlayerIndex(game.players[0]) {
			player1Count++
		}
		if game.board[startRow][startCol] == game.getPlayerIndex(game.players[1]) {
			player2Count++
		}
		startRow += stepRow
		startCol += stepCol
	}
	return player1Count, player2Count
}

func (game *Game) getPlayerIndex(player string) int {
	if game.players[0] == player {
		return 1
	} else if game.players[1] == player {
		return 2
	} else {
		// TODO: flag error
		return -1
	}
}

func (game *Game) getNextPlayer(curPlayer string) string {
	// assume exactly 2 players
	for _, val := range game.players {
		if val != curPlayer {
			return val
		}
	}

	// TODO: throw error
	return "UnknownPlayer"
}
