package cli

import (
	// "log"
	// "fmt"
	// "fmt"
    "os"
    "os/exec"
    "runtime"
    // "time"
)

var opt_args = make(map[string]string)
var clear map[string]func() //create a map for storing clear funcs

func Parse(args_list []string, callback func(map[string]interface{})) {
	clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }

	main_args, doujin_code := "", ""

	if(len(args_list) > 1){
		main_args = args_list[1]
	}

	if(len(args_list) > 2) {
		doujin_code = args_list[2]
		
		for i, v := range args_list {
			if i > 2 {
				opt_args[v] = v
			}
		}
	} else {
		opt_args = map[string]string{}
	}

	parsed := map[string]interface{}{
		"main_args" 	: main_args,
		"doujin_code"   : doujin_code,
		"opt_args"		: opt_args,
	}

	// fmt.Println(opt_args)
	callback(parsed)
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}