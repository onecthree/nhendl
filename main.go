package main

import (
	"os"
	// "fmt"
	"github.com/onecthree/nhendl/cli"
	"github.com/onecthree/nhendl/http"
	"github.com/onecthree/nhendl/error"
	"github.com/onecthree/nhendl/handle"
	"github.com/onecthree/nhendl/tags"
	"github.com/onecthree/nhendl/download"
	// "runtime"
)

const nh = "https://nhentai.net/g/"

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())

	cli.Parse(os.Args, func(data map[string]interface{}) {
		error.Start(data, map[string]func(){
			"direct" : func() {
				handle.Run(http.Get(nh + data["doujin_code"].(string)), func(src string, code int) {
					error.Parser(code)
		
					doujin_tags := tags.Default(src)
					doujin_tags.Source = src
					doujin_tags.Clear()
					
					doujin_tags.Preview(data["opt_args"].(map[string]string))
		
					download.Prompt(data["doujin_code"].(string), doujin_tags.Pages(), doujin_tags.Host(data["doujin_code"].(string)), data["opt_args"].(map[string]string))
					
				})
			},
		})
	})


}