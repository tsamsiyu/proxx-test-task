package pkg

import "errors"

var GameOverErr = errors.New("game is lost")

type Gameplay struct {
	gameField *GameField
}

func NewGameplay(gameField *GameField) *Gameplay {
	return &Gameplay{
		gameField: gameField,
	}
}

func (gp *Gameplay) Open(row int, col int) error {
	cell, err := gp.gameField.Get(row, col)
	if err != nil {
		return err
	}

	if cell.IsBlackHole {
		return GameOverErr
	}

	cell.IsOpened = true

	if cell.AdjacentBlackHoles == 0 {
		if err := gp.openAdjacentCells(cell); err != nil {
			return err
		}
	}

	return nil
}

func (gp *Gameplay) openAdjacentCells(cell *Cell) error {
	boundary := cell.AdjacentBoundary

	for adjacentRow := boundary.TopLeft.Row; adjacentRow <= boundary.BottomRight.Row; adjacentRow++ {
		for adjacentCol := boundary.TopLeft.Col; adjacentCol <= boundary.BottomRight.Col; adjacentCol++ {
			adjacentCell, err := gp.gameField.Get(adjacentRow, adjacentCol)
			if err != nil {
				return err
			}
			if adjacentCell.IsOpened {
				continue
			}

			if err := gp.Open(adjacentRow, adjacentCol); err != nil {
				return err
			}
		}
	}

	return nil
}
