package error

import (
	"fmt"
	"os"
	"regexp"
)

func code_check(code string, data map[string]interface{}) {
	is_code, _ := regexp.MatchString("^[0-9]{1,9}$", code)
	if(!is_code) {
		fmt.Println("nhendl: invalid doujin code for argument -- \"" + data["doujin_code"].(string) +"\"\n\nSee `nhendl -h` for more options.")
		os.Exit(0)
	}
}

func opt_exists(x string) bool {
	opt_list := []string{
		"--async",
		"--fulldesc",
	}

	error := true
	for _, v := range opt_list {
		if x == v {
			error = false
		}
	}

	if error {
		return false
	}

	return true
}

func Start(data map[string]interface{}, callback map[string]func()) {
	if(len(data["opt_args"].(map[string]string)) > 0) {
		for _, v := range data["opt_args"].(map[string]string) {
			if !opt_exists(v) {
				fmt.Println("nhendl: invalid optional arguments -- \""+ v  +"\"\n\nSee `nhendl -h` for more options.")
				os.Exit(0)
			}
		} 
	}

	// declare regex checker for first arguments if was actual doujin code

	switch data["main_args"].(string) {
		case "-d":
			code_check(data["doujin_code"].(string), data)
			fmt.Println("nhendl: searching for -- \"https://nhentai.net/g/" + data["doujin_code"].(string) + "\"")
			callback["direct"]()	
		case "-h":
			fmt.Print("nhendl 1.0, a light-weight nhentai doujin downloader.\n")
			fmt.Print("Usage: nhendl -d [CODE] [OPTION]...")
			fmt.Print("\n\n")
			fmt.Print("Startup:\n")
			fmt.Print("   -v,                 display the version of nhendl.\n")
			fmt.Print("   -h,                 print this help.\n")

			fmt.Print("\nOption:\n")
			fmt.Print("   --fulldesc,         print preview of doujin info as full description.\n")
			fmt.Print("   --async,            use goroutines to spawn each downloader for all pages,\n")
			fmt.Print("                       require a fast connection for significant result.")
			fmt.Print("\n")
			os.Exit(0)
		default:
			fmt.Println("nhendl: invalid option -- \""+ data["main_args"].(string) +"\"\nUsage: nhendl -d [CODE]... [OPTION]...\n\nSee `nhendl -h` for more options.")
			os.Exit(0)
	}

}

func Parser(code int) {
	if(code == 404) {
		fmt.Println("\nnhendl: doujin could not be found. [404]")
		os.Exit(0)
	}
}