package controllers

import "strconv"
import "runtime"

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

func getCallStackText() string {
	const callStackTextLengthLimit = 64 * 1024
	var callStackTextData = make([]byte, callStackTextLengthLimit)
	var callStackTextLength = runtime.Stack(callStackTextData, true)
	var callStackText = string(callStackTextData[:callStackTextLength])
	return callStackText
}

func exceptionToText(e error) string {
	return e.Error() + "\n" + getCallStackText()
}
