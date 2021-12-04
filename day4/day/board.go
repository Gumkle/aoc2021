package day

type Board [Side][Side]uint8

func (b *Board) WinsBy(takenShots Shots) bool {
	for i := 0; i < Side; i++ {
		if takenShots.Contain(b[i][i]) {
			rowWins, colWins := true, true
			for j := 0; j < Side; j++ {
				if !takenShots.Contain(b[i][j]) {
					rowWins = false
				}
				if !takenShots.Contain(b[j][i]) {
					colWins = false
				}
				if !rowWins && !colWins {
					break
				}
			}
			if rowWins || colWins {
				return true
			}
		}
	}
	return false
}

func (b *Board) SumOfUnmarkedBy(takenShots Shots) (sum uint64) {
	for i := 0; i < Side; i++ {
		for j := 0; j < Side; j++ {
			if !takenShots.Contain(b[i][j]) {
				sum += uint64(b[i][j])
			}
		}
	}
	return
}
