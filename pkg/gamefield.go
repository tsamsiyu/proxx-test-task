package pkg

import (
	"errors"
	"math/rand"
)

// GameField
// The data structure for storing game state is openAdjacentCells 2-dimensional array
// The purpose for choosing it corresponds to the access pattern of the game and provides quick memory access (log(row*col))
type GameField struct {
	Width  int
	Height int
	Cells  [][]Cell
}

type CellPosition struct {
	Row int
	Col int
}

// Cell Represents the state of openAdjacentCells game field unit and stores all its necessary metadata
type Cell struct {
	Pos                CellPosition
	AdjacentBoundary   CellsBoundary
	IsOpened           bool
	IsBlackHole        bool
	AdjacentBlackHoles int
}

type CellsBoundary struct {
	TopLeft     CellPosition
	BottomRight CellPosition
}

func (gf *GameField) Get(row int, col int) (*Cell, error) {
	if row > gf.Height-1 || row < 0 {
		return nil, errors.New("out of game field's boundary")
	}
	if col > gf.Width-1 || col < 0 {
		return nil, errors.New("out of game field's boundary")
	}

	return &gf.Cells[row][col], nil
}

func (gf *GameField) IsAllOpened() bool {
	for _, row := range gf.Cells {
		for _, cell := range row {
			if !cell.IsOpened && !cell.IsBlackHole {
				return false
			}
		}
	}

	return true
}

func GenerateGameField(height int, width int, blackHoles int) (*GameField, error) {
	if height < 3 || width < 3 {
		return nil, errors.New("field's size can't be less than 3x3")
	}

	if blackHoles > height*width/2 {
		return nil, errors.New("not allowed to set black holes in more than half of all cells")
	}

	cells := initCells(height, width)
	fillBlackHoles(cells, height, width, blackHoles)
	fillAdjacentBlackHoles(cells)

	return &GameField{
		Width:  width,
		Height: height,
		Cells:  cells,
	}, nil
}

func initCells(height int, width int) [][]Cell {
	cells := make([][]Cell, height)

	for row := 0; row < height; row++ {
		cells[row] = make([]Cell, width)

		for col := 0; col < width; col++ {
			pos := CellPosition{
				Row: row,
				Col: col,
			}

			cells[row][col] = Cell{
				Pos:              pos,
				AdjacentBoundary: buildAdjacentCellsBoundary(&pos, height, width),
			}
		}
	}

	return cells
}

// fillBlackHoles
func fillBlackHoles(cells [][]Cell, height int, width int, blackHoles int) {
	var generatedBlackHoles int

	for generatedBlackHoles < blackHoles {
		row := rand.Intn(height - 1)
		col := rand.Intn(width - 1)

		cell := &cells[row][col]

		if !cell.IsBlackHole {
			cell.IsBlackHole = true
			generatedBlackHoles++
		}
	}
}

// fillAdjacentBlackHoles
func fillAdjacentBlackHoles(cells [][]Cell) {
	for ri := range cells {
		for ci := range cells[ri] {
			cell := &cells[ri][ci]
			if cell.IsBlackHole {
				continue
			}

			cell.AdjacentBlackHoles = computeBlackHolesInBoundaries(&cell.AdjacentBoundary, cells)
		}
	}
}

func buildAdjacentCellsBoundary(pos *CellPosition, height int, width int) CellsBoundary {
	topLeft := CellPosition{
		Row: pos.Row - 1,
		Col: pos.Col - 1,
	}

	bottomRight := CellPosition{
		Row: pos.Row + 1,
		Col: pos.Col + 1,
	}

	if topLeft.Row < 0 {
		topLeft.Row = 0
	}

	if topLeft.Col < 0 {
		topLeft.Col = 0
	}

	if bottomRight.Row > height-1 {
		bottomRight.Row = height - 1
	}

	if bottomRight.Col > width-1 {
		bottomRight.Col = width - 1
	}

	return CellsBoundary{
		TopLeft:     topLeft,
		BottomRight: bottomRight,
	}
}

func computeBlackHolesInBoundaries(boundary *CellsBoundary, cells [][]Cell) int {
	var adjacentHoles int

	for row := boundary.TopLeft.Row; row <= boundary.BottomRight.Row; row++ {
		for col := boundary.TopLeft.Col; col <= boundary.BottomRight.Col; col++ {
			if cells[row][col].IsBlackHole {
				adjacentHoles++
			}
		}
	}

	return adjacentHoles
}
