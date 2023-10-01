package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

const (
	ScreenWidth    = 1000
	ScreenHeight   = 600
	BoardSize      = 9
	CellSize       = 66
	ButtonWidth    = 120
	ButtonHeight   = 40
	ButtonX        = ScreenWidth - ButtonWidth - 20
	ButtonY        = 10
	ResultX        = 10
	ResultY        = 70
	ResultTextSize = 24
)

var (
	board               [BoardSize][BoardSize]int
	selectedRow         = -1
	selectedCol         = -1
	verificationInProgess = false
)

func Starting_game() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Sudoku Example")
	rl.SetTargetFPS(60)

	rand.Seed(time.Now().UnixNano())

	// Generate a valid Sudoku grid with some empty cells
	generateSudoku()

	for !rl.WindowShouldClose() {
		handleInput()

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			mouseX := rl.GetMouseX()
			mouseY := rl.GetMouseY()

			// Check if the "Check" button was clicked
			if mouseX >= ButtonX && mouseX <= ButtonX+ButtonWidth && mouseY >= ButtonY && mouseY <= ButtonY+ButtonHeight {
				verificationInProgess = true
				result := checkGrid()
				fmt.Println("The grid is correct:", result)
			}

			// Check if the "Generate" button was clicked
			if mouseX >= ButtonX && mouseX <= ButtonX+ButtonWidth && mouseY >= ButtonY+50 && mouseY <= ButtonY+50+ButtonHeight {
				generateSudoku()
				verificationInProgess = false
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Draw the Sudoku grid
		drawGrid()
		drawNumbers()

		// Draw the "Check" button
		rl.DrawRectangle(int32(ButtonX), int32(ButtonY), int32(ButtonWidth), int32(ButtonHeight), rl.RayWhite)
		rl.DrawRectangleLines(int32(ButtonX), int32(ButtonY), int32(ButtonWidth), int32(ButtonHeight), rl.Black)
		rl.DrawText("Vérifier", int32(ButtonX+10), int32(ButtonY+10), 20, rl.Black)

		// Draw the "Generate" button
		rl.DrawRectangle(int32(ButtonX), int32(ButtonY+50), int32(ButtonWidth), int32(ButtonHeight), rl.RayWhite)
		rl.DrawRectangleLines(int32(ButtonX), int32(ButtonY+50), int32(ButtonWidth), int32(ButtonHeight), rl.Black)
		rl.DrawText("Généré", int32(ButtonX+10), int32(ButtonY+60), 20, rl.Black)

		// Display a verification result message
		if verificationInProgess {
			message := "The grid is correct!"
			if !checkGrid() {
				message = "The grid is incorrect."
			}
			rl.DrawText(message, int32(ResultX), int32(ResultY), int32(ResultTextSize), rl.Black)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// Check if the Sudoku grid is correct.
func checkGrid() bool {
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			num := board[row][col]
			if num != 0 {
				for i := 0; i < BoardSize; i++ {
					if i != col && board[row][i] == num {
						return false
					}
					if i != row && board[i][col] == num {
						return false
					}
				}
				startRow, startCol := row-row%3, col-col%3
				for i := startRow; i < startRow+3; i++ {
					for j := startCol; j < startCol+3; j++ {
						if i != row && j != col && board[i][j] == num {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

// Generate a valid Sudoku grid with some empty cells.
func generateSudoku() {
	// Clear the grid
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			board[row][col] = 0
		}
	}

	// Fill the grid following Sudoku rules
	solveSudoku()

	// Limit the number of randomly generated digits
	emptyCells := 45 // Adjust this number according to your preference
	for emptyCells > 0 {
		row := rand.Intn(BoardSize)
		col := rand.Intn(BoardSize)

		if board[row][col] != 0 {
			board[row][col] = 0
			emptyCells--
		}
	}
}

// Solve the Sudoku grid.
func solveSudoku() bool {
	row, col := findEmptyCell()
	if row == -1 && col == -1 {
		return true // Sudoku solved
	}

	for num := 1; num <= BoardSize; num++ {
		if isSafe(row, col, num) {
			board[row][col] = num

			if solveSudoku() {
				return true
			}

			board[row][col] = 0 // If no solution is found, backtrack
		}
	}

	return false
}

// Check if a number can be safely placed in a cell.
func isSafe(row, col, num int) bool {
	// Check the row
	for i := 0; i < BoardSize; i++ {
		if board[row][i] == num {
			return false
		}
	}

	// Check the column
	for i := 0; i < BoardSize; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// Check the 3x3 region
	startRow, startCol := row-row%3, col-col%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}

// Find the first empty cell in the grid.
func findEmptyCell() (int, int) {
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if board[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1 // No empty cell found
}

// Handle player input.
func handleInput() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mouseX := float32(rl.GetMouseX())
		mouseY := float32(rl.GetMouseY())

		selectedRow = int(mouseY / CellSize)
		selectedCol = int(mouseX / CellSize)
	}

	for key := rl.KeyOne; key <= rl.KeyNine; key++ {
		if rl.IsKeyPressed(int32(key)) && selectedRow != -1 && selectedCol != -1 {
			// Allow the user to modify any cell, even if it's already filled
			board[selectedRow][selectedCol] = int(key - rl.KeyOne + 1)
		}
	}

	if rl.IsKeyPressed(rl.KeyR) {
		generateSudoku() // Generate a new random grid
		verificationInProgess = false
	}
}

// Draw the Sudoku grid.
func drawGrid() {
	for i := 0; i <= BoardSize; i++ {
		lineThickness := float32(1)
		if i%3 == 0 {
			lineThickness = float32(3) // Lignes plus épaisses pour délimiter les régions
		}

		// Draw horizontal and vertical lines of the grid
		rl.DrawLineEx(
			rl.NewVector2(0, float32(i)*CellSize),
			rl.NewVector2(float32(BoardSize)*CellSize, float32(i)*CellSize),
			lineThickness,
			rl.Black,
		)
		rl.DrawLineEx(
			rl.NewVector2(float32(i)*CellSize, 0),
			rl.NewVector2(float32(i)*CellSize, float32(BoardSize)*CellSize),
			lineThickness,
			rl.Black,
		)
	}
}

// Draw the numbers in the grid.
func drawNumbers() {
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			num := board[row][col]
			if num != 0 {
				// Draw the number in the center of the cell
				x := int32(float32(col)*CellSize + CellSize/2 - 10)
				y := int32(float32(row)*CellSize + CellSize/2 - 10)
				rl.DrawText(fmt.Sprintf("%d", num), x, y, 20, rl.Black)
			}

			// Highlight the selected cell
			if row == selectedRow && col == selectedCol {
				rl.DrawRectangleLines(
					int32(float32(col)*CellSize),
					int32(float32(row)*CellSize),
					int32(CellSize),
					int32(CellSize),
					rl.Red,
				)
			}
		}
	}
}
