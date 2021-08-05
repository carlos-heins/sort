package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	piece := createpiece(200000)
	quicksort(piece)

	var stringval (string)
	for i := 0; i < len(piece); i++ {
		stringval += strconv.Itoa(piece[i])
		stringval += " "
	}
	saveFile(stringval, "sorted")
}
func saveFile(file string, sorted string) {
	message := []byte(file)
	err := ioutil.WriteFile("./storages/"+sorted+".txt", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func createpiece(size int) []int {
	piece := make([]int, size, size)
	var stringvalues (string)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		piece[i] = rand.Intn(200000) - rand.Intn(200000)
		stringvalues += strconv.Itoa(piece[i])
		stringvalues += " "
	}
	saveFile(stringvalues, "unsorted")
	return piece
}
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	center := rand.Int() % len(a)

	a[center], a[right] = a[right], a[center]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
