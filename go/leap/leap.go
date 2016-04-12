package leap

func IsLeapYear(year int) bool {
	return divisibleBy(4, year) && (!divisibleBy(100, year) || divisibleBy(400, year))
}

func divisibleBy(denominator int, numerator int) bool {
	return (numerator % denominator) == 0
}