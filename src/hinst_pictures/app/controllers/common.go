package controllers

import "strconv"

func formatInteger(x int) string {
	var text = strconv.Itoa(x)
	var formattedText = ""
	for index, character := range text {
		var reversedIndex = len(text) - index - 1
		formattedText = formattedText + string(character)
		if reversedIndex%3 == 0 {
			formattedText = formattedText + " "
		}
	}
	return formattedText
}
