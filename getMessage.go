package main

import (
	"strings"
)

func deleteElementFromArray(array []string, index int) []string {
	return append(array[:index], array[index+1:]...)
}

func deleteEmptyWord(words []string) []string {
    var res []string
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
		return words,deleted,positions
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


func GetMessage(satellites []SatelliteRequest) string {
	var index = []int {0,0,0,}
	var message = []string {}

	for true {
		var candidateWord = []string {}

		for i:=0; i < len(index); i++ {
			if index[i] < len(satellites[i].Message) {
				candidateWord = append(candidateWord, satellites[i].Message[index[i]])	
			}
		}

		if len(candidateWord) == 0 {
			break
		}
	
		candidateWord, deleted, positions := deletePreviousWord(candidateWord, message)
		if deleted {
			for _, p := range positions {
				index[p]++
			}
			continue
		}
 		candidateWord = deleteEmptyWord(candidateWord)
		candidateWord = deleteDuplicatedWord(candidateWord)

		if len(candidateWord) == 1 {
			message = append(message, candidateWord[0])
		}

		for i := 0; i < len(index); i++ {
			index[i]++
		}

	}

	return strings.Join(message[:], " ")
}