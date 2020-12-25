package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	size = 8
	popSize = 150
	mutationProb = 3
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

func removeFromSlice(s Board, i int) Board {
	for k, v := range s{
		if v == i{
			s[k], s[len(s)-1] = s[len(s)-1], s[k]
			break
		}
	}
	return s[:len(s)-1]
}

func getConflict(x1 int, y1 int, x2 int, y2 int) bool {
	if (x2 - x1 == Abs(y2 - y1)) || y1 == y2 {
		return true
	}
	return false
}

func getFitness(board Board) int {
	// board for size=8 --> (4, 0, 3, 5, 7, 1, 6, 2)
	fitness := 0
	firstConflict := true
	for x1, y1 := range board {
		firstConflict = true
		for i := x1 + 1; i < size; i++ {
			if getConflict(x1, y1, i, board[i]) == true {
				fitness++
				if firstConflict == true {
					fitness++
					firstConflict = false
				}
			}
		}
	}

	return fitness
}

func populationCombSort(population Population) Population {
	alen := len(population)
	gap := alen * 10 / 13
	swapped := true

	for true {
		if 8 < gap && gap < 11 {
			gap = 11
		}
		swapped = false
		for i := 0; i < alen - gap; i++ {
			if getFitness(population[i + gap]) < getFitness(population[i]) {
				population[i], population[i + gap] = population[i + gap], population[i]
				swapped = true
			}
		}
		if gap * 10 / 13 > 0 {
			gap = gap * 10 / 13
		} else if !swapped {
			break
		}
	}

	return population
}

func crossover(parent1 Board, parent2 Board) Board {
	descendant := Board{}
	indexes := Board{}
	storage := rand.Perm(size)
	randInt := 0

	for i := 0; i < size; i++ {
		if parent1[i] == parent2[i] {
			descendant = append(descendant, parent1[i])
			storage = removeFromSlice(storage, parent1[i])
		} else {
			descendant = append(descendant, 0)
			indexes = append(indexes, i)
		}
	}

	for _, i := range indexes{
		randInt = rand.Intn(len(storage))
		descendant[i] = storage[randInt]
		storage = removeFromSlice(storage, storage[randInt])
	}

	return descendant
}

func gemmation(board Board) Board {
	i1 := rand.Intn(size)
	i2 := rand.Intn(size)

	board[i1], board[i2] = board[i2], board[i1]

	return board
}

func mutation(board Board) Board {
	return gemmation(board)
}

func runCrossover(population Population) Population {
	for i := 1; i < popSize / 2; i++{
		descendant := crossover(population[i-1], population[i])
		//descendant := gemmation(population[i])
		if mutationProb / 100 < rand.Float32() {
			descendant = mutation(descendant)
		}
		population[popSize / 2 + i] = descendant
	}

	return population
}

func initPopulation() Population {
	population := Population{}
	for i := 0; i < popSize; i++ {
		population[i] = rand.Perm(size)
	}
	return population
}

func visualize(board Board)  {
	for _, i := range board {
		for j := 0; j < size; j++ {
			if i == j {
				fmt.Print("Q  ")
			} else {
				fmt.Print("+  ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	population := initPopulation()
	population = populationCombSort(population)
	if getFitness(population[0]) == 0 {
		visualize(population[0])
	} else {
		for true {
			population = runCrossover(population)
			population = populationCombSort(population)
			if getFitness(population[0]) == 0 {
				visualize(population[0])
				fmt.Println(getFitness(population[0]))
				break
			}
		}
	}
}
