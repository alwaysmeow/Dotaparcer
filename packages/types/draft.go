package types

import "fmt"

type Draft struct {
	Heroes   [5]*Hero
	Accuracy float64
	Winrate  float64
}

type DraftGrid [5][5]float64

func (draft *Draft) Log() {
	for _, hero := range draft.Heroes {
		hero.Log()
	}
	fmt.Printf("Accuracy: %.2f\n", draft.Accuracy)
	fmt.Printf("Winrate: %.2f\n", draft.Winrate)
}

func (draft *Draft) Error() float64 {
	grid := CreateDraftGrid(draft.Heroes)
	err := 0.
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			err += grid[i][j] * grid[j][i]
		}
	}
	return err
}

func CreateDraft(heroes [5]*Hero) Draft {
	grid := CreateDraftGrid(heroes)

	accuracy := 1.

	for !grid.solved() {
		rounded := grid.fixation()
		accuracy *= rounded
		grid.correction()
	}

	draft := Draft{Accuracy: accuracy}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j] == 1 {
				draft.Heroes[j] = heroes[i]
			}
		}
	}

	winrate := 0.
	for i := 0; i < 5; i++ {
		winrate += float64(draft.Heroes[i].Winrate[i])
	}
	draft.Winrate = winrate / 5

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

func (grid *DraftGrid) fixation() float64 {
	x, y := -1, -1

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j] < 1 {
				if x+y < 0 || grid[i][j] > grid[x][y] {
					x, y = i, j
				}
			} else {
				for k := 0; k < 5; k++ {
					if k != i && grid[k][j] > 0 {
						x, y = i, j
					}
				}
			}
		}
	}

	rounded := grid[x][y]

	for i := 0; i < 5; i++ {
		grid[x][i] = 0
		grid[i][y] = 0
	}
	grid[x][y] = 1

	return rounded
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
