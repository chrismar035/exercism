package romannumerals

import (
	"errors"
	"sort"
	"strconv"
	"sync"
)

const testVersion = 2

type InPacket struct {
	value int
	order int
}

type OutPacket struct {
	value string
	order int
}

func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 || arabic > 3999 {
		return "", errors.New("Roman numerals must be between 1 and 3999")
	}

	var wg sync.WaitGroup
	wire := make(chan OutPacket)
	arabics := strconv.Itoa(arabic)

	for i, char := range arabics {
		wg.Add(1)
		value, _ := strconv.Atoi(string(char))
		go func(output chan OutPacket, in InPacket) {
			defer wg.Done()

			output <- OutPacket{
				value: placeString(mags[in.order], in.value),
				order: in.order}

		}(wire, InPacket{order: len(arabics) - i - 1, value: value})
	}
	go func() {
		defer close(wire)
		wg.Wait()
	}()

	results := collectPackets(wire)
	sort.Sort(ByOrder(results))
	return concatValues(results), nil
}

func concatValues(sl []OutPacket) string {
	result := ""
	for _, out := range sl {
		result = out.value + result
	}

	return result
}

func collectPackets(wire chan OutPacket) []OutPacket {
	results := []OutPacket{}
	for out := range wire {
		results = append(results, out)
	}
	return results
}
