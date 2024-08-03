package types

import (
	"fmt"
)

type Draft struct {
	Heroes   [5]*Hero
	Accuracy float64
	Winrate  float64
	Meta     float64
}

type DraftGrid [5][5]float64

func (draft *Draft) Log() {
	for _, hero := range draft.Heroes {
		hero.Log()
	}
	fmt.Printf("Accuracy: %.2f\n", draft.Accuracy)
	fmt.Printf("Winrate: %.2f\n", draft.Winrate)
	fmt.Printf("Meta: %.4f\n", draft.Meta)
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
	meta := 0.
	for i := 0; i < 5; i++ {
		winrate += float64(draft.Heroes[i].Winrate[i])
		meta += float64(draft.Heroes[i].Meta[i])
	}
	draft.Winrate = winrate / 5
	draft.Meta = meta

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

	maxMetrica := 0.

	for pos := 0; pos < 5; pos++ {
		sum := 0.
		for hero := 0; hero < 5; hero++ {
			sum += grid[hero][pos]
		}
		if sum > 0 {
			for hero := 0; hero < 5; hero++ {
				monopoly := grid[hero][pos] / sum
				metrica := monopoly * grid[hero][pos]
				if metrica < 1 {
					if monopoly == 1 {
						x, y = hero, pos
						maxMetrica = 1
						break
					}
					if x+y < 0 || metrica > maxMetrica {
						x, y = hero, pos
						maxMetrica = metrica
					}
				}
			}
		}
	}

	for i := 0; i < 5; i++ {
		grid[x][i] = 0
		grid[i][y] = 0
	}
	grid[x][y] = 1

	return maxMetrica
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
