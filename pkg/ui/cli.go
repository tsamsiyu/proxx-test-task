package ui

import (
	"fmt"

	"proxx-test-task/pkg"
)

type Input struct {
	Height     int
	Width      int
	BlackHoles int
}

func RunCliView() error {
	input, err := PromptCliInput()
	if err != nil {
		return err
	}

	fmt.Println()

	gameField, err := pkg.GenerateGameField(input.Height, input.Width, input.BlackHoles)
	if err != nil {
		return err
	}

	DrawCliGameField(gameField)

	gameplay := pkg.NewGameplay(gameField)

	if err := InteractiveCliPlay(gameField, gameplay); err != nil {
		return err
	}

	return nil
}

func DrawCliGameField(gameField *pkg.GameField) {
	fmt.Print("    ")
	for col := 0; col < gameField.Width; col++ {
		fmt.Printf("%d ", col)
	}

	for row := 0; row < gameField.Height; row++ {
		fmt.Printf("\n")
		fmt.Printf("%d   ", row)

		for col := 0; col < gameField.Width; col++ {
			cell := gameField.Cells[row][col]
			if cell.IsOpened {
				if cell.IsBlackHole {
					fmt.Print("X ", col)
				} else {
					fmt.Printf("%d ", cell.AdjacentBlackHoles)
				}
			} else {
				fmt.Print("? ")
			}
		}
	}

	fmt.Printf("\n")
}

func InteractiveCliPlay(gameField *pkg.GameField, gameplay *pkg.Gameplay) error {
	for {
		fmt.Println()

		row, err := IntPrompt("Open row #:")
		if err != nil {
			return err
		}

		col, err := IntPrompt("Open column #:")
		if err != nil {
			return err
		}

		if err := gameplay.Open(row, col); err != nil {
			if err == pkg.GameOverErr {
				fmt.Println("Game is lost")
				return nil
			} else {
				return err
			}
		}

		fmt.Println()

		if gameField.IsAllOpened() {
			fmt.Println("Game is won")
			return nil
		}

		DrawCliGameField(gameField)
	}
}

func PromptCliInput() (*Input, error) {
	height, err := IntPrompt("Height:")
	if err != nil {
		return nil, err
	}

	width, err := IntPrompt("Width:")
	if err != nil {
		return nil, err
	}

	blackHoles, err := IntPrompt("Black holes:")
	if err != nil {
		return nil, err
	}

	return &Input{
		Height:     height,
		Width:      width,
		BlackHoles: blackHoles,
	}, nil
}
