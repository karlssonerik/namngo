package names

import (
	"math/rand/v2"
)

func FilterAnimalNamesByLength(maxLength int) []string {
	var filteredNames []string
	for _, name := range AnimalNames {
		if len(name) <= maxLength {
			filteredNames = append(filteredNames, name)
		}
	}

	return filteredNames
}

func GetRandomAnimalName(names []string) string {
	if len(names) == 0 {
		return ""
	}
	index := rand.IntN(len(names))
	return names[index]
}
