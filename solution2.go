package main

import (
	"fmt"
	"math"
	"strings"
)

func solution2() {
	fmt.Println(decode("LLRR="))
	fmt.Println(decode("==RLL"))
	fmt.Println(decode("=LLRR"))
	fmt.Println(decode("=RRLL"))
	fmt.Println(decode("RRL=R"))
	fmt.Println(decode("RRRL="))
}

func decode(encoded string) string {
	result := []int{}

	switch encoded[0] {
	case 'L':
		result = append(result, 1, 0)
	case 'R':
		result = append(result, 0, 1)
	default:
		result = append(result, 0, 0)
	}

	// fmt.Println(result)

	for i := 1; i < len(encoded); i++ {
		code := encoded[i]

		if code == 'L' {
			toBeAppended := 0
			firstL := -1

			for j := len(result) - 1; j >= 1; j-- {
				if result[j] <= result[j-1] {
					if math.Abs(float64(result[j]-result[j-1])) > 1 {
						firstL = j
					} else {
						firstL = j - 1
					}
				} else if result[j] > result[j-1] {
					break
				} else {
					toBeAppended = 1
				}
			}

			if firstL > -1 {
				for k := firstL; k < len(result); k++ {
					result[k] += 1
				}
			}

			result = append(result, toBeAppended)

		} else if code == 'R' {

			for j := len(result) - 1; j >= 1; j-- {
				if result[j] < result[j-1] {
					break
				}
			}

			result = append(result, result[len(result)-1]+1)
		} else if code == '=' {
			for j := len(result) - 1; j >= 1; j-- {
				if result[j] < result[j-1] || result[j] > result[j-1] {
					break
				} else {
					result[j-1] = result[j]
				}
			}

			result = append(result, result[len(result)-1])
		}

	}

	return strings.Join(strings.Split(fmt.Sprint(result), " "), "")
}
