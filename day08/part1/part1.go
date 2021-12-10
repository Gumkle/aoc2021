package part1

var SimpleDigitMap = map[int]int{
	2: 1,
	3: 7,
	7: 8,
	4: 4,
}

func CountDigits(outputs [][]string) uint {
	vector := getCountVector([]int8{1, 4, 7, 8})
	counts := [10]uint{}
	for _, set := range outputs {
		for _, digit := range set {
			if digit == "" {
				continue
			}
			parsedDigit := parse(digit)
			counts[parsedDigit] += uint(vector[parsedDigit])
		}
	}
	sum := uint(0)
	for _, value := range counts {
		sum += value
	}
	return sum
}

func getCountVector(u []int8) [10]int8 {
	result := [10]int8{}
	for _, value := range u {
		result[value] = 1
	}
	return result
}

func parse(digit string) int {
	value, ok := SimpleDigitMap[len(digit)]
	if ok {
		return value
	}
	return 0
}
