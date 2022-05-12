package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
	// "github.com/briandowns/spinner"
)

var current = 1

func spinner() {
	chara := map[string]string{
		"1" : "|",
		"2" : "/",
		"3" : "-",
		"4" : "\\",
	}

	for {
		if current > 4 {
			current = 1
		}
		fmt.Println(chara[strconv.Itoa(current)])
		current++
		time.Sleep(100 * time.Millisecond)		
	}
}

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

func loading(pages int, current int) string {
	var x = 20.0 / float64(pages)
	var r = math.Round(x * float64(current))
	var percent = 100.0 / float64(pages)
	result := loop(int(r))
	result += " (" + strconv.Itoa(current) + "/" + strconv.Itoa(pages) + ") " + strconv.Itoa(int(percent * float64(current))) + "%"
	return result
}

func main() {
	// s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	// s.Prefix = loading(27, 26) + " "
	// s.Start()
	// time.Sleep(4 * time.Second)
	spinner()
}