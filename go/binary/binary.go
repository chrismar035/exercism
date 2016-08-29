package binary

import (
	"errors"
	"math"
	"strings"
	"sync"
)

// ParseBinary will convert a binary number, represented as a string
// (e.g. '101010'), to its decimal equivalent using first principles
func ParseBinary(raw string) (int, error) {
	var wg sync.WaitGroup
	ch := make(chan int)
	length := len(raw)
	chars := strings.Split(raw, "")

	for i, char := range chars {
		switch char {
		case "1":
			wg.Add(1)
			go computeDecimalPlace(ch, length-i-1, &wg)
		case "0":
			continue
		default:
			return 0, errors.New("String contains non-binary characters")
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	sum := 0
	for i := range ch {
		sum += i
	}
	return sum, nil
}

func computeDecimalPlace(ch chan int, place int, wg *sync.WaitGroup) {
	ch <- int(math.Exp2(float64(place)))
	wg.Done()
}
