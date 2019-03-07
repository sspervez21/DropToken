package droptoken

import "98point6DropToken/models"

// Movetype enum
type Movetype int

const (
	move Movetype = 0
	quit Movetype = 1
)

// Move in a drop token game
type Move struct {
	moveNumber int
	mType      Movetype
	player     string
	column     int
}

func convertMove(m Move) *models.Move {
	moveStr := "MOVE"
	quitStr := "QUIT"

	var model *models.Move

	if m.mType == move {
		model = &models.Move{Type: &moveStr, Column: int64(m.column), Player: &m.player}
	} else if m.mType == quit {
		model = &models.Move{Type: &quitStr, Player: &m.player}
	} else {
		// TODO: signal error
	}

	return model
}

func convertMoves(moves []Move) []*models.Move {
	returnedMoves := []*models.Move{}
	for _, m := range moves {
		returnedMoves = append(returnedMoves, convertMove(m))
	}

	return returnedMoves
}
