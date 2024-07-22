package types

import (
	"fmt"
	"math"
)

type Draft [5]*Hero
type DraftGrid [5][5]float64

func (draft Draft) log() {
	for i := 0; i < 5; i++ {
		draft[i].log()
	}
	fmt.Printf("Accuracy: %.2f\n", draft.Accuracy())
	fmt.Printf("Winrate: %.2f\n", draft.Winrate())
}

func CreateDraft(heroes [5]*Hero) Draft {
	grid := CreateDraftGrid(heroes)

	for !grid.solved() {
		grid.fixation()
		grid.correction()
	}

	var draft Draft

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j] == 1 {
				draft[j] = heroes[i]
			}
		}
	}

	return draft
}

func CreateDraftGrid(heroes [5]*Hero) DraftGrid {
	var grid DraftGrid

	for heroIndex, hero := range heroes {
		matches := 0
		for _, posMatches := range hero.Matches {
			matches += posMatches
		}
		for posIndex, posMatches := range hero.Matches {
			grid[heroIndex][posIndex] = float64(posMatches) / float64(matches)
		}
	}

	return grid
}

func (grid *DraftGrid) solved() bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j] != 0 && grid[i][j] != 1 {
				return false
			}
		}
	}
	return true
}

func (grid *DraftGrid) fixation() {
	x, y := -1, -1

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j] < 1 {
				if x+y < 0 || grid[i][j] > grid[x][y] {
					x, y = i, j
				}
			}
		}
	}
	for i := 0; i < 5; i++ {
		grid[x][i] = 0
		grid[i][y] = 0
	}
	grid[x][y] = 1
}

func (grid *DraftGrid) correction() {
	for row := 0; row < 5; row++ {
		sum := 0.
		for col := 0; col < 5; col++ {
			sum += grid[row][col]
		}
		for col := 0; col < 5; col++ {
			grid[row][col] = grid[row][col] / sum
		}
	}
}

func (draft *Draft) Accuracy() float64 {
	grid := CreateDraftGrid(*draft)
	acc := 1.
	for i := 0; i < 5; i++ {
		acc *= float64(grid[i][i])
	}
	acc = math.Pow(acc, 1.0/5.0)
	return acc
}

func (draft *Draft) Winrate() float64 {
	acc := 1.
	for i := 0; i < 5; i++ {
		acc *= float64(draft[i].Winrate[i])
	}
	acc = math.Pow(acc, 1.0/5.0)
	return acc
}
