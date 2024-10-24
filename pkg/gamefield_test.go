package pkg

import "testing"

func TestGenerateGameField(t *testing.T) {
	t.Run("Number of generated black holes corresponds to the expected one", func(t *testing.T) {
		gameField, err := GenerateGameField(4, 4, 1)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		actualBlackHoles := 0
		for _, row := range gameField.Cells {
			for _, cell := range row {
				if cell.IsBlackHole {
					actualBlackHoles++
				}
			}
		}

		if actualBlackHoles != 1 {
			t.Errorf("Actual black holes number mismatches initial assignment: %d", actualBlackHoles)
		}
	})
}

func TestFillAdjacentBlackHoles(t *testing.T) {
	t.Run("Number of adjacent black holes is calculated correctly for one hole in the center", func(t *testing.T) {
		cells := initCells(5, 5)
		cells[2][2].IsBlackHole = true

		fillAdjacentBlackHoles(cells)

		expectedAdjacencies := [][]int{
			{0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{0, 1, -1, 1, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0},
		}

		for expectationRow, expectationCols := range expectedAdjacencies {
			for expectationCol, expectationValue := range expectationCols {
				if expectationValue == -1 {
					continue
				}

				actualCell := cells[expectationRow][expectationCol]

				if actualCell.AdjacentBlackHoles != expectationValue {
					t.Errorf(
						"calculated adjacent black holes at position %d:%d is wrong, expected: %d, got: %d",
						expectationRow,
						expectationCol,
						expectationValue,
						actualCell.AdjacentBlackHoles,
					)
				}
			}
		}
	})

	t.Run("Number of adjacent black holes is calculated correctly", func(t *testing.T) {
		cells := initCells(4, 4)
		cells[0][0].IsBlackHole = true
		cells[2][2].IsBlackHole = true
		cells[3][3].IsBlackHole = true

		fillAdjacentBlackHoles(cells)

		expectedAdjacentValues := [][]int{
			{-1, 1, 0, 0},
			{1, 2, 1, 1},
			{0, 1, -1, 2},
			{0, 1, 2, -1},
		}

		for expectationRow, expectedRowValue := range expectedAdjacentValues {
			for expectationCol, expectationValue := range expectedRowValue {
				if expectationValue == -1 {
					continue
				}

				actualCell := cells[expectationRow][expectationCol]
				if actualCell.AdjacentBlackHoles != expectationValue {
					t.Errorf(
						"calculated adjacent black holes at position %d:%d is wrong, expected: %d, got: %d",
						expectationRow,
						expectationCol,
						expectationValue,
						actualCell.AdjacentBlackHoles,
					)
				}
			}
		}
	})
}
