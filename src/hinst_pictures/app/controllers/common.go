package controllers

import "strconv"

func formatInteger(x int) string {
	var text = strconv.Itoa(x)
	return text
}
