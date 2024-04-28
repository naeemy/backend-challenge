package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sync"
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
	var cache sync.Map

	for w := 0; w < 1_000_000; w++ {
		go worker(arr, jobs, &result, &cache)
	}

	firstPrev := ""
	secondPrev := ""

	half := new(big.Int).Div(totalPaths, big.NewInt(2))
	for path := new(big.Int).Set(big.NewInt(0)); path.Cmp(half) < 0; path.Add(path, big.NewInt(1)) {
		var directions string
		directions = createJob(arr, path)
		jobs <- []string{directions, firstPrev}
		firstPrev = directions

		secondHalf := new(big.Int).Add(half, path)
		directions = createJob(arr, secondHalf)
		jobs <- []string{directions, secondPrev}
		secondPrev = directions
	}

	close(jobs)

	fmt.Println(result)
}

func createJob(arr [][]int, path *big.Int) string {
	leading := new(big.Int).Lsh(big.NewInt(1), uint(len(arr)))
	directionInt := new(big.Int).Add(path, leading)
	directions := fmt.Sprintf("%b", directionInt)
	directions = directions[1:]

	return directions
}

func worker(arr [][]int, jobs <-chan []string, max *int, cache *sync.Map) {

	for job := range jobs {
		sum := 0
		i := 0
		j := 0
		directions := job[0]
		lastPath := job[1]

		for k := 0; k < len(directions); k++ {

			if lastPath == "" {
				break
			}

			if directions[k] != lastPath[k] {
				key := createCacheKey(k-1, j)
				cached, _ := cache.Load(key)
				if cached != nil {
					i = k
					sum = cached.(int)

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

		for i < len(directions) {
			if directions[i] == '0' {
				sum += arr[i][j]
			} else {
				j++
				sum += arr[i][j]
			}

			if i != len(directions)-1 {
				key := createCacheKey(i, j)
				cache.Store(key, sum)
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
