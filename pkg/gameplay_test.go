package pkg

import (
	"strconv"
	"testing"
)

func TestGameplay_Open(t *testing.T) {
	t.Run("Simulated 5-steps game", func(t *testing.T) {
		cells := initCells(6, 6)
		cells[0][5].IsBlackHole = true
		cells[4][4].IsBlackHole = true
		cells[5][3].IsBlackHole = true
		fillAdjacentBlackHoles(cells)

		//{"0", "0", "0", "0", "1", "X"},
		//{"0", "0", "0", "0", "1", "1"},
		//{"0", "0", "0", "0", "0", "0"},
		//{"0", "0", "0", "1", "1", "1"},
		//{"0", "0", "1", "2", "X", "1"},
		//{"0", "0", "1", "X", "2", "1"},

		gameField := GameField{
			Height: 6,
			Width:  6,
			Cells:  cells,
		}

		gamePlay := Gameplay{
			gameField: &gameField,
		}

		// OPEN 1

		if err := gamePlay.Open(5, 2); err != nil {
			t.Error(err)
		}

		expectedState := [][]string{
			{"?", "?", "?", "?", "?", "?"},
			{"?", "?", "?", "?", "?", "?"},
			{"?", "?", "?", "?", "?", "?"},
			{"?", "?", "?", "?", "?", "?"},
			{"?", "?", "?", "?", "?", "?"},
			{"?", "?", "1", "?", "?", "?"},
		}

		assertExpectedState(t, expectedState, cells)

		// open 2

		if err := gamePlay.Open(0, 0); err != nil {
			t.Error(err)
		}

		expectedState = [][]string{
			{"0", "0", "0", "0", "1", "?"},
			{"0", "0", "0", "0", "1", "1"},
			{"0", "0", "0", "0", "0", "0"},
			{"0", "0", "0", "1", "1", "1"},
			{"0", "0", "1", "2", "?", "?"},
			{"0", "0", "1", "?", "?", "?"},
		}

		assertExpectedState(t, expectedState, cells)

		// open 3

		if err := gamePlay.Open(5, 4); err != nil {
			t.Error(err)
		}

		expectedState = [][]string{
			{"0", "0", "0", "0", "1", "?"},
			{"0", "0", "0", "0", "1", "1"},
			{"0", "0", "0", "0", "0", "0"},
			{"0", "0", "0", "1", "1", "1"},
			{"0", "0", "1", "2", "?", "?"},
			{"0", "0", "1", "?", "2", "?"},
		}

		assertExpectedState(t, expectedState, cells)

		// open 4

		if err := gamePlay.Open(5, 5); err != nil {
			t.Error(err)
		}

		expectedState = [][]string{
			{"0", "0", "0", "0", "1", "?"},
			{"0", "0", "0", "0", "1", "1"},
			{"0", "0", "0", "0", "0", "0"},
			{"0", "0", "0", "1", "1", "1"},
			{"0", "0", "1", "2", "?", "?"},
			{"0", "0", "1", "?", "2", "1"},
		}

		assertExpectedState(t, expectedState, cells)

		// open 5

		if err := gamePlay.Open(4, 5); err != nil {
			t.Error(err)
		}

		expectedState = [][]string{
			{"0", "0", "0", "0", "1", "?"},
			{"0", "0", "0", "0", "1", "1"},
			{"0", "0", "0", "0", "0", "0"},
			{"0", "0", "0", "1", "1", "1"},
			{"0", "0", "1", "2", "?", "1"},
			{"0", "0", "1", "?", "2", "1"},
		}

		assertExpectedState(t, expectedState, cells)
	})
}

func assertExpectedState(t *testing.T, expectedState [][]string, cells [][]Cell) {
	for ri, row := range cells {
		for ci, cell := range row {
			expectedValStr := expectedState[ri][ci]

			if expectedValStr == "?" {
				if cell.IsOpened {
					t.Errorf("Cell at %d:%d expected to be closed", ri, ci)
				}
				continue
			}

			if expectedValStr == "X" && !cell.IsBlackHole {
				t.Errorf("Cell at %d:%d expected to be openAdjacentCells black hole", ri, ci)
				continue
			}

			expectedValInt, _ := strconv.ParseInt(expectedValStr, 10, 64)

			if expectedValInt != int64(cell.AdjacentBlackHoles) {
				t.Errorf(
					"Cell at %d:%d has mismatched adjacent black holes number, %d expected, %d given",
					ri,
					ci,
					expectedValInt,
					cell.AdjacentBlackHoles,
				)
				continue
			}
		}
	}
}
