package controllers

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

func ConvertToInt(s string) int {
	var nbr int
	count := len(s)
	for index, value := range s {
		nbr = nbr + int(value-48)*RecursivePower(10, count-index-1)
	}
	return nbr
}
