package leap

func IsLeapYear(year int) bool {
	switch {
	case !divisibleBy(4, year):
		return false
	case divisibleBy(100, year) && !divisibleBy(400, year):
		return false
	}
	return true
}

func divisibleBy(denominator int, numerator int) bool {
	return (numerator % denominator) == 0
}

