package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
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

	jobs := make(chan []string)
	result := 0
	cache := make(map[string]int)

	for w := 0; w < 1; w++ {
		go worker(arr, jobs, &result, &cache)
	}

	lastPath := ""

	for path := new(big.Int).Set(big.NewInt(0)); path.Cmp(totalPaths) < 0; path.Add(path, big.NewInt(1)) {
		leading := new(big.Int).Lsh(big.NewInt(1), uint(len(arr)))
		directionInt := new(big.Int).Add(path, leading)
		directions := fmt.Sprintf("%b", directionInt)
		directions = directions[1:]

		jobs <- []string{directions, lastPath}

		lastPath = directions
	}

	close(jobs)

	fmt.Println(result)
}

func worker(arr [][]int, jobs <-chan []string, max *int, cache *map[string]int) {

	for job := range jobs {
		// fmt.Println()
		sum := 0
		i := 0
		j := 0
		directions := job[0]
		lastPath := job[1]

		// fmt.Println(directions, lastPath)
		for k := 0; k < len(directions); k++ {

			if lastPath == "" {
				break
			}

			if directions[k] != lastPath[k] {
				key := createCacheKey(k-1, j)
				// fmt.Println(key, sum)
				if (*cache)[key] != 0 {
					i = k
					sum = (*cache)[key]

					break
				} else {
					i = 0
					j = 0
					break
				}

			}

			if directions[k] == '1' {
				j++
			}

		}

		// fmt.Println(*cache, i, j)
		for i < len(directions) {

			// fmt.Println(i, j)

			if directions[i] == '0' {
				sum += arr[i][j]
			} else {
				j++
				sum += arr[i][j]
			}

			if i != len(directions)-1 {
				key := createCacheKey(i, j)

				(*cache)[key] = sum
			}

			i++
		}

		if sum > *max {
			*max = sum
			fmt.Println(*max)
		}
	}

}

func createCacheKey(i int, j int) string {
	return fmt.Sprintf("%d:%d", i, j)
}
