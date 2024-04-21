package main

import (
	"fmt"
	"math"
)

func solution2() {
	fmt.Println(decode("LLRR="))
	fmt.Println(decode("==RLL"))
	fmt.Println(decode("=LLRR"))
	fmt.Println(decode("RRL=R"))
	fmt.Println(decode("RRRL="))
}

func findMin(list []int) int {
	minimum := math.MaxInt64

	for _, number := range list {
		if number <= minimum {
			minimum = number
		}
	}

	return minimum
}

func decode(encoded string) []int {
	cursor := []int{}
	current := 0
	const CHANGE = 1

	// init
	firstChar := encoded[0]
	if firstChar == 'L' {
		cursor = append(cursor, CHANGE, current)
	} else if firstChar == 'R' {
		cursor = append(cursor, current, CHANGE)
		current += CHANGE
	} else {
		cursor = append(cursor, current, current)
	}

	// draw
	for i := 1; i < len(encoded); i++ {
		if encoded[i] == 'L' {
			if encoded[i-1] == 'R' {
				current = 0
			} else {
				current -= CHANGE
			}
			cursor = append(cursor, current)
		} else if encoded[i] == 'R' {
			current += CHANGE
			cursor = append(cursor, current)
		} else {
			cursor = append(cursor, current)
		}
	}

	adjust(encoded, &cursor)
	return cursor
}

func adjust(encoded string, cursor *[]int) {

	// increase
	toAdjust := int(math.Abs(float64(findMin(*cursor) - 0)))
	for i := 0; i < len(encoded); i++ {
		if i == 1 {
			if encoded[i-1] == '=' {
				(*cursor)[0] = (*cursor)[i]
			} else if encoded[i-1] == 'L' {
				(*cursor)[0] = (*cursor)[i] + 1
			}
		}

		(*cursor)[i+1] += toAdjust
	}

	// check starting with '='
	if encoded[0] == '=' {
		temp := []int{0, 0}
		for i := 1; i < len(encoded); i++ {
			if encoded[i] == '=' {
				temp = append(temp, 0)
			} else if encoded[i] == 'L' && len(temp) > 2 {
				if sum(temp) == 0 {
					for j := range temp {
						temp[j] = (*cursor)[i+1]
					}
				}
			} else if encoded[i] == 'R' && len(temp) > 2 {
				if sum(temp) == 0 {
					for j := range temp {
						temp[j] = 0
					}
				}
			}

			if len(temp) > 2 && (encoded[i] == 'R' || encoded[i] == 'L') {
				newCursor := append(temp, (*cursor)[len(temp):]...)
				copy(*cursor, newCursor)
				break
			}

		}
	}
}

func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
