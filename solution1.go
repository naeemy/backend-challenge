package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func solution1() {

	// Uncomment to test
	// arr := [][]int{
	// 	{59},
	// 	{73, 41},
	// 	{52, 40, 53},
	// 	{26, 53, 6, 34},
	// }

	arr := [][]int{}
	data, _ := os.ReadFile("./files/hard.json")
	json.Unmarshal(data, &arr)

	n, e := big.NewInt(2), big.NewInt(int64(len(arr)-1))

	n.Exp(n, e, nil)

	maxChan := make(chan int)
	finished := make(chan bool)
	result := make(chan int)

	go compare(maxChan, finished, result)

	for path := new(big.Int).Set(big.NewInt(0)); path.Cmp(n) < 0; path.Add(path, big.NewInt(1)) {
		leading := new(big.Int).Lsh(big.NewInt(1), uint(len(arr)))
		directionInt := new(big.Int).Add(path, leading)
		directions := fmt.Sprintf("%b", directionInt)
		directions = directions[1:]

		go findMax(arr, directions, maxChan)
		finished <- false
	}

	close(maxChan)
	finished <- true
	close(finished)
	fmt.Println(<-result)
	close(result)

}

func findMax(arr [][]int, directions string, value chan<- int) {
	sum := 0
	j := 0

	for i, char := range directions {

		direction, _ := strconv.Atoi(string(char))

		if direction == 0 {
			sum += arr[i][j]
		} else {
			j++
			sum += arr[i][j]
		}
	}

	value <- sum
}

func compare(value <-chan int, finished <-chan bool, result chan<- int) {

	maxValue := 0
	isFinished := false
	for !isFinished {
		received := <-value
		isFinished = <-finished
		if received > maxValue {
			maxValue = received
			fmt.Println(maxValue)
		}
	}

	result <- maxValue

}
