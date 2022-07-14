package Codewars

/*
RGB To Hex Conversion
The rgb function is incomplete. Complete it so that passing in RGB decimal values will result in a hexadecimal
representation being returned. Valid decimal values for RGB are 0 - 255.
Any values that fall out of that range must be rounded to the closest valid value.

The following are examples of expected output values:

kata.rgb(255, 255, 255) -- returns FFFFFF
kata.rgb(255, 255, 300) -- returns FFFFFF
kata.rgb(0, 0, 0) -- returns 000000
kata.rgb(148, 0, 211) -- returns 9400D3
*/

import (
	"fmt"
	"strings"
)

func RGB(r, g, b int) string {
	colors := [3]int{r, g, b}
	for idx, elem := range colors {
		if elem < 0 {
			colors[idx] = 0
		}
		if elem > 255 {
			colors[idx] = 255
		}
	}
	s:= fmt.Sprintf("%02x%02x%02x", colors[0], colors[1], colors[2])
	return strings.ToUpper(s)
}