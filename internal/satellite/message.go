// Package satellite stores struct and methods which handle satellites information
package satellite

import (
	"strings"
)

func deleteElementFromArray(array []string, index int) (res []string) {
	if len(array) == 0 {
		return array
	}
	res = append(array[:index], array[index+1:]...)
	return
}

func deleteEmptyWord(words []string) []string {
	var res []string

	if len(words) == 0 {
		return words
	}

	for _, str := range words {
		if str != "" {
			res = append(res, str)
		}
	}
	return res
}

func deletePreviousWord(words []string, partialMessage []string) ([]string, bool, []int) {
	positions := []int{}
	var deleted = false

	if len(partialMessage) == 0 {
		return words, deleted, positions
	}

	for index := 0; index < len(words); index++ {
		if words[index] == partialMessage[len(partialMessage)-1] {
			words = deleteElementFromArray(words, index)
			positions = append(positions, index)
			deleted = true
		}
	}

	return words, deleted, positions
}

func deleteDuplicatedWord(words []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range words {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// GetMessage receives data from all satellites and tries to decrypts the ships message
func GetMessage(satellites []Data) string {
	var index = []int{0, 0, 0}
	var message = []string{}

	for true {

		// populate candidateWord which contains each satellite word of n-th iteration
		var candidateWord = []string{}
		for i := 0; i < len(index); i++ {
			if index[i] < len(satellites[i].Message) {
				candidateWord = append(candidateWord, satellites[i].Message[index[i]])
			}
		}
		if len(candidateWord) == 0 {
			break
		}

		// if a word is the same as previous read then ignore it and retry the iteration
		candidateWord, deleted, positions := deletePreviousWord(candidateWord, message)
		if deleted {
			for _, p := range positions {
				index[p]++
			}
			continue
		}
		candidateWord = deleteEmptyWord(candidateWord)
		candidateWord = deleteDuplicatedWord(candidateWord)

		// in case of candidateWord has 1 element it means that this element is the word!
		if len(candidateWord) == 1 {
			message = append(message, candidateWord[0])
		}

		for i := 0; i < len(index); i++ {
			index[i]++
		}

	}

	return strings.Join(message[:], " ")
}
