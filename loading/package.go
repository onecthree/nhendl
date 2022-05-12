package loading

import (
	"math"
	"strconv"
)

func loop(lim int) string {
	bar := "[ "
	i := 0;
	for i < 20 {
		if i < lim {
			bar += "="
		} else {
			bar += "-"
		}
		i++
	}
	bar += " ]"
	return bar
}

func Loading(pages int, current int) string {
	var x = 20.0 / float64(pages)
	var r = math.Round(x * float64(current))
	var percent = 100.0 / float64(pages)
	result := loop(int(r))
	result += " (" + strconv.Itoa(current) + "/" + strconv.Itoa(pages) + ") " + strconv.Itoa(int(percent * float64(current))) + "%"
	return result
}