package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

const (
	screenWidth  = 600
	screenHeight = 600
	gridSize     = 9
)

var (
	grid        [gridSize][gridSize]int
	cellSize    float32
	selectedRow = -1
	selectedCol = -1
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Sudoku Example")
	rl.SetTargetFPS(60)

	cellSize = float32(screenWidth) / float32(gridSize)

	// Initialiser le générateur de nombres aléatoires avec une graine basée sur le temps
	rand.Seed(time.Now().UnixNano())

	// Générer une grille Sudoku valide avec quelques cases vides
	generateSudoku()

	for !rl.WindowShouldClose() {
		handleInput()
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		drawGrid()
		drawNumbers()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func generateSudoku() {
	// Effacer la grille
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			grid[row][col] = 0
		}
	}

	// Remplir la grille en respectant les règles du Sudoku
	solveSudoku()

	// Limiter le nombre de chiffres générés aléatoirement
	remaining := 81 - countInitialNumbers()
	for remaining > 0 {
		row := rand.Intn(gridSize)
		col := rand.Intn(gridSize)

		if grid[row][col] == 0 {
			num := rand.Intn(9) + 1
			if isSafe(row, col, num) {
				grid[row][col] = num
				remaining--
			}
		}
	}

	// Supprimer certaines valeurs pour créer des cases vides
	emptyCells := 45 // Nombre de cases vides (ajustez-le selon vos préférences)
	for emptyCells > 0 {
		row := rand.Intn(gridSize)
		col := rand.Intn(gridSize)

		if grid[row][col] != 0 {
			grid[row][col] = 0
			emptyCells--
		}
	}
}

func solveSudoku() bool {
	emptyRow, emptyCol := findEmptyCell() // Ignorer la deuxième valeur

	if emptyRow == -1 && emptyCol == -1 {
		return true // La grille est résolue
	}

	for num := 1; num <= 9; num++ {
		if isSafe(emptyRow, emptyCol, num) {
			grid[emptyRow][emptyCol] = num

			if solveSudoku() {
				return true
			}

			grid[emptyRow][emptyCol] = 0 // Annuler la tentative
		}
	}

	return false // Aucune solution trouvée pour cette configuration
}

func isSafe(row, col, num int) bool {
	// Vérifier la ligne
	for i := 0; i < gridSize; i++ {
		if grid[row][i] == num {
			return false
		}
	}

	// Vérifier la colonne
	for i := 0; i < gridSize; i++ {
		if grid[i][col] == num {
			return false
		}
	}

	// Vérifier la région 3x3
	startRow, startCol := row-row%3, col-col%3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if grid[i][j] == num {
				return false
			}
		}
	}

	return true
}

func findEmptyCell() (int, int) {
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if grid[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1 // Aucune cellule vide trouvée
}

func countInitialNumbers() int {
	count := 0
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if grid[row][col] != 0 {
				count++
			}
		}
	}
	return count
}

func handleInput() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mouseX := float32(rl.GetMouseX())
		mouseY := float32(rl.GetMouseY())

		selectedRow = int(mouseY / cellSize)
		selectedCol = int(mouseX / cellSize)
	}

	if rl.IsKeyPressed(rl.KeySpace) && selectedRow != -1 && selectedCol != -1 {
		// Ne rien faire si l'utilisateur tente de modifier un chiffre généré automatiquement
		if grid[selectedRow][selectedCol] == 0 {
			grid[selectedRow][selectedCol] = -1 // Marquer la cellule comme modifiée par l'utilisateur
		}
	}

	for key := rl.KeyOne; key <= rl.KeyNine; key++ {
		if rl.IsKeyPressed(int32(key)) && selectedRow != -1 && selectedCol != -1 {
			// Ajouter le chiffre sélectionné à la grille si la cellule est modifiable
			if grid[selectedRow][selectedCol] == -1 {
				grid[selectedRow][selectedCol] = int(key - rl.KeyOne + 1)
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyBackspace) && selectedRow != -1 && selectedCol != -1 {
		// Effacer la cellule sélectionnée si elle est modifiable
		if grid[selectedRow][selectedCol] == -1 {
			grid[selectedRow][selectedCol] = 0
		}
	}
}

func drawGrid() {
	for i := 0; i < gridSize; i++ {
		// Dessiner les lignes horizontales et verticales
		rl.DrawLine(int32(float32(i)*cellSize), 0, int32(float32(i)*cellSize), screenHeight, rl.Black)
		rl.DrawLine(0, int32(float32(i)*cellSize), screenWidth, int32(float32(i)*cellSize), rl.Black)

		// Dessiner des lignes épaisses pour délimiter les régions 3x3
		if i%3 == 0 && i != 0 {
			rl.DrawLine(int32(float32(i)*cellSize), 0, int32(float32(i)*cellSize), screenHeight, rl.Black)
			rl.DrawLine(0, int32(float32(i)*cellSize), screenWidth, int32(float32(i)*cellSize), rl.Black)
		}
	}
}

func drawNumbers() {
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			number := grid[row][col]
			if number != 0 && number != -1 {
				// Dessiner le numéro au centre de la case
				x := int32(float32(col)*cellSize + cellSize/2 - 10)
				y := int32(float32(row)*cellSize + cellSize/2 - 10)
				rl.DrawText(fmt.Sprintf("%d", number), x, y, 20, rl.Black)
			}

			// Dessiner la case sélectionnée en surbrillance
			if row == selectedRow && col == selectedCol {
				rl.DrawRectangleLines(int32(float32(col)*cellSize), int32(float32(row)*cellSize), int32(cellSize), int32(cellSize), rl.Red)
			}
		}
	}
}
