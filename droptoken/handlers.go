package droptoken

import (
	"98point6DropToken/models"
	"98point6DropToken/restapi/operations"
	"sort"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
)

// CreateGameHandler creates a new drop token game
func CreateGameHandler(params operations.CreateGameParams) middleware.Responder {
	// TODO: validate that we were provided exactly 2 players
	if len(params.NewGame.Players) != 2 || *params.NewGame.Rows != 4 || *params.NewGame.Columns != 4 {
		return operations.NewCreateGameBadRequest().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.CreateGameBadRequestCode),
			Message: "This game is designed for 2 players on a 4x4 board.",
		})
	}

	game := createGame([2]string{params.NewGame.Players[0], params.NewGame.Players[1]}, int(*params.NewGame.Columns), int(*params.NewGame.Rows))
	return operations.NewCreateGameOK().WithPayload(&models.Game{GameID: game.gameID})
}

// GetAllGamesHandler returns all active games
func GetAllGamesHandler(params operations.GetAllGamesParams) middleware.Responder {
	games := getAllGames()
	sort.Sort(byGameID(games))
	return operations.NewGetAllGamesOK().WithPayload(&models.AllGames{Games: games})
}

// GetGameHandler returns the state of the given gameId
func GetGameHandler(params operations.GetGameParams) middleware.Responder {
	game, ok := getGame(params.GameID)

	if !ok {
		return operations.NewGetGameNotFound().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.GetGameNotFoundCode),
			Message: params.GameID + " not found.",
		})
	}

	gameState := &models.GameState{}

	gameState.Players = []string{game.players[0], game.players[1]}

	inProgressState := "IN_PROGRESS"
	doneState := "DONE"
	errorState := "ERROR"

	if game.state == inProgress {
		gameState.State = &inProgressState
	} else if game.state == done {
		gameState.State = &doneState

		// TODO: ensure 'winner' is in the players list
		gameState.Winner = game.winner
	} else {
		// TODO: indicate internal server error
		gameState.State = &errorState
	}

	return operations.NewGetGameOK().WithPayload(gameState)
}

// GetAllMovesHandler returns all the moves for the given game
func GetAllMovesHandler(params operations.GetAllMovesParams) middleware.Responder {
	game, ok := getGame(params.GameID)

	if !ok {
		return operations.NewGetAllMovesNotFound().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.GetAllMovesNotFoundCode),
			Message: params.GameID + " not found.",
		})
	}

	var moves = []Move{}

	for _, m := range game.moves {
		moves = append(moves, m)
	}

	sort.Sort(byMoveNumber(moves))
	return operations.NewGetAllMovesOK().WithPayload(&models.AllMoves{Moves: convertMoves(moves)})
}

// MakeMoveHandler makes the specified move
func MakeMoveHandler(params operations.MakeMoveParams) middleware.Responder {
	game, ok := getGame(params.GameID)

	if !ok {
		return operations.NewMakeMoveNotFound().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.MakeMoveNotFoundCode), // 404
			Message: params.GameID + " not found.",
		})
	}

	if game.state == done {
		return operations.NewMakeMoveBadRequest().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.MakeMoveBadRequestCode), // 400
			Message: "Game over.",
		})
	}

	if !game.isPlayerInGame(params.PlayerID) {
		return operations.NewMakeMoveNotFound().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.MakeMoveNotFoundCode), // 404
			Message: params.PlayerID + " not part of this game.",
		})
	}

	// not this player's turn
	if params.PlayerID != game.curPlayer {
		return operations.NewMakeMoveConflict().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.MakeMoveConflictCode), // 409
			Message: params.PlayerID + " played out of turn.",
		})
	}

	if !game.isLegalMove(params.PlayerID, int(*params.NewMove.Column)) {
		return operations.NewMakeMoveBadRequest().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.MakeMoveBadRequestCode), // 400
			Message: "Illegal move.",
		})
	}

	moveNumber := game.makeMove(params.PlayerID, int(*params.NewMove.Column))

	moveOutputStr := game.gameID + "/moves/" + strconv.Itoa(moveNumber)
	return operations.NewMakeMoveOK().WithPayload(&models.MoveOutput{Move: &moveOutputStr})
}

// GetMoveHandler makes the specified move
func GetMoveHandler(params operations.GetMoveParams) middleware.Responder {
	game, ok := getGame(params.GameID)

	if !ok {
		return operations.NewGetMoveNotFound().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.GetMoveNotFoundCode), // 404
			Message: params.GameID + " not found.",
		})
	}

	m, ok := game.moves[int(params.MoveNumber)]

	if !ok {
		return operations.NewGetMoveNotFound().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.GetMoveNotFoundCode), // 404
			Message: strconv.Itoa(int(params.MoveNumber)) + " not found.",
		})
	}

	return operations.NewGetMoveOK().WithPayload(convertMove(m))
}

// RemovePlayerHandler makes the specified move
func RemovePlayerHandler(params operations.RemovePlayerParams) middleware.Responder {
	game, ok := getGame(params.GameID)

	if !ok {
		return operations.NewRemovePlayerNotFound().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.RemovePlayerNotFoundCode), // 404
			Message: params.GameID + " not found.",
		})
	}

	if !game.isPlayerInGame(params.PlayerID) {
		return operations.NewRemovePlayerNotFound().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.RemovePlayerNotFoundCode), // 404
			Message: params.PlayerID + " not part of this game.",
		})
	}

	if game.state == done {
		return operations.NewRemovePlayerGone().WithPayload(&models.MalformedRequest{
			Code:    int64(operations.RemovePlayerGoneCode), // 410
			Message: "Game already over.",
		})
	}

	// remove player from game
	game.removePlayer(params.PlayerID)

	return operations.NewRemovePlayerAccepted()
}
