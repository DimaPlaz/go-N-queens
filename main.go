package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	size = 10
	popSize = 5000
)

type (
	Board []int
	Population [popSize]Board
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getConflict(x1 int, y1 int, x2 int, y2 int) bool {
	if x2 - x1 == Abs(y2 - y1) || y1 == y2{
		return true
	}
	return false
}

func getFitness(board Board) int {
	// for size=8 --> (4, 0, 3, 5, 7, 1, 6, 2)
	fitness := 0

	for x1, y1 := range board{
		for i := x1 + 1; i < size; i++{
			if getConflict(x1, y1, i, board[i]) == true{
				fitness++
				continue
			}
		}
	}

	return fitness
}

func populationCombSort(population Population) Population {
	alen := len(population)
	gap := alen * 10 / 13
	swapped := true

	for true{
		if 8 < gap && gap < 11{
			gap = 11
		}
		swapped = false
		for i := 0; i < alen - gap; i++{
			if getFitness(population[i + gap]) < getFitness(population[i]) {
				population[i], population[i + gap] = population[i + gap], population[i]
				swapped = true
			}
		}
		if gap * 10 / 13 > 0{
			gap = gap * 10 / 13
		} else if !swapped {
			break
		}
	}

	return population
}

func initPopulation() Population {
	population := Population{}
	for i := 0; i < popSize; i++{
		population[i] = rand.Perm(size)
	}
	return population
}

func visualize(board Board)  {
	for _, i := range board{
		for j := 0; j < size; j++{
			if i == j{
				fmt.Print("Q ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Print("\n")
	}
}

func main()  {
	rand.Seed(time.Now().UnixNano())

	population := initPopulation()
	population = populationCombSort(population)
	visualize(population[0])
	fmt.Println(getFitness(population[0]))

}
