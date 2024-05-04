package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"runtime"
	"sync"
)

func solution1() {
	wg := new(sync.WaitGroup)
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
	totalRoutes := base.Exp(base, power, nil)
	chunkSize := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil)

	jobs := make(chan []string)
	result := 0
	var cache sync.Map

	for w := 0; w < runtime.NumCPU(); w++ {
		go worker(arr, jobs, &result, &cache)
	}

	var incrementBy big.Int

	if totalRoutes.Cmp(chunkSize) >= 0 {
		incrementBy = *chunkSize
	} else {
		incrementBy = *new(big.Int).Sub(totalRoutes, big.NewInt(1))
	}

	var start big.Int = *big.NewInt(0)
	var stop big.Int = *big.NewInt(0)

	// TODO: FIX RACE CONDITION
	for task := incrementBy; task.Cmp(totalRoutes) < 0; task.Add(&task, &incrementBy) {
		wg.Add(1)
		stop = task
		go func(s big.Int, e big.Int) {
			lastRoute := ""
			for route := &s; route.Cmp(&e) < 0 || route.Cmp(&e) == 0; route.Add(route, big.NewInt(1)) {
				directions := createJob(arr, *route)
				jobs <- []string{directions, lastRoute}
				lastRoute = directions

			}
			wg.Done()
		}(start, stop)
		start = task
	}

	wg.Wait()
	close(jobs)

	fmt.Println(result)
}

func generateChunks(incrementBy big.Int, totalRoutes *big.Int, chunks chan<- []big.Int) {

}
func createJob(arr [][]int, route big.Int) string {
	leading := new(big.Int).Lsh(big.NewInt(1), uint(len(arr)))
	directionInt := new(big.Int).Add(&route, leading)
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
