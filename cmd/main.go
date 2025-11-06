package main

import (
	"flag"
	"fmt"
	"math/rand/v2"

	"github.com/karlssonerik/namngo/internal/names"
)

func main() {
	length := flag.Int("l", 10, "maximum length of animal names to filter")
	flag.Parse()

	filteredNames := names.FilterAnimalNamesByLength(*length)
	randomName := names.GetRandomAnimalName(filteredNames)
	missing := *length - len(randomName)
	for missing >= 3 {
		newNames := names.FilterAnimalNamesByLength(missing)
		randomName += names.GetRandomAnimalName(newNames)
		missing = missing - len(randomName)
	}

	for missing > 0 {
		randomName += fmt.Sprintf("%d", rand.IntN(9))
		missing -= 1
	}

	println(randomName)
}
