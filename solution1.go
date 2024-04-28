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

	base, power := big.NewInt(2), big.NewInt(int64(len(arr)-1))

	totalPaths := base.Exp(base, power, nil)

	jobs := make(chan string)
	result := 0

	for w := 0; w < 1_000_000; w++ {
		go worker(arr, jobs, &result)
	}

	for path := new(big.Int).Set(big.NewInt(0)); path.Cmp(totalPaths) < 0; path.Add(path, big.NewInt(1)) {
		leading := new(big.Int).Lsh(big.NewInt(1), uint(len(arr)))
		directionInt := new(big.Int).Add(path, leading)
		directions := fmt.Sprintf("%b", directionInt)
		directions = directions[1:]

		jobs <- directions
	}

	close(jobs)

	fmt.Println(result)
}

func worker(arr [][]int, jobs <-chan string, max *int) {

	for directions := range jobs {
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

		if sum > *max {
			*max = sum
			fmt.Println(*max)
		}
	}

}
