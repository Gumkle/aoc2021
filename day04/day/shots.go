package day

type Shots []uint8

func (s *Shots) Contain(number uint8) bool {
	for _, val := range *s {
		if val == number {
			return true
		}
	}
	return false
}
