package download

import (
	"log"
	"os"
	"bufio"
	"fmt"
	"time"
	// "strconv"
	"github.com/onecthree/nhendl/http"
	// "github.com/onecthree/nhendl/cli"
	"github.com/onecthree/nhendl/loading"
	"github.com/briandowns/spinner"
	"sync"
)

var current_page = 0
var pages = 0
var s = spinner.New(spinner.CharSets[9], 100*time.Millisecond)

func syncProcied(code string, host string) {
	s.Prefix = "nhendl: doujin (#"+code+"): " + loading.Loading(pages, current_page) + " "
	s.Suffix = " [sync] (Downloading)"
	s.Start()

	start := time.Now()

	syncProcessDownload(code, host)
	
	time.Sleep(100 * time.Millisecond)
	fmt.Print(" (Finished ")
	fmt.Print(time.Since(start))
	fmt.Print(")")
	fmt.Print("\n")
}

var limit_conc = 5

func increaseLoad(pagex int, lc int) {
	calc := pagex - lc
	if(calc > 3) {
		limit_conc += 3
	} else {
		limit_conc = pages
	}
}

func Prompt(code string, page int, host string, opt_args map[string]string) {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("nhendl: continue to download doujin? (Y/N): ")
	text, _ := reader.ReadString('\n')

	switch text {
		case "Y\n", "y\n":
			if err := os.Mkdir(code, os.ModePerm); err != nil {
				log.Fatal(err)
			}	
			pages = page

			if(len(opt_args) > 0) {
				for _, v := range opt_args {
					switch v {
						case "--async":
							// increaseLoad(pages, limit_conc)
							// fmt.Println(limit_conc)
							
							s.Prefix = "nhendl: doujin (#"+code+"): " + loading.Loading(pages, current_page) + " "
							s.Suffix = " [async] (Downloading)"
							s.Start()

							start := time.Now()
							asyncProcessDownload(code, host)
							
							time.Sleep(100 * time.Millisecond)
							fmt.Print(" (Finished ")
							fmt.Print(time.Since(start))
							fmt.Print(")")
							fmt.Print("\n")
						case "--fulldesc":
							
						default:
							syncProcied(code, host)
					}
				}
			} else {
				syncProcied(code, host)
			}

		case "N\n", "n\n":
			os.Exit(0)
	}
	
	os.Exit(0)
}

func syncProcessDownload(code string, host string) {
	s.Prefix = "nhendl: doujin (#"+code+"): " + loading.Loading(pages, current_page) + " "
	http.Download(code, current_page, host, func() {
		current_page++
	})

	if current_page <= pages {
		syncProcessDownload(code, host)
	} else {
		s.Suffix = " [sync] (Completed)"
	}
}

var async_page = 0
var mu = &sync.Mutex{}
var wg sync.WaitGroup

func spawnThread(code string, cat int, host string) {
	wg.Add(1)
	http.Download(code, cat, host, func() {
		defer wg.Done()

		mu.Lock()
		current_page++
		async_page++
		s.Prefix = "nhendl: doujin (#"+code+"): " + loading.Loading(pages, async_page) + " "
		mu.Unlock()

		if(current_page <= pages) {
			spawnThread(code, current_page, host)
		}
	})
}

func asyncProcessDownload(code string, host string) {

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(cat int) {
			defer wg.Done()
			spawnThread(code, cat, host)
			// http.Download(code, cat, host, func() {
			// 	defer wg.Done()
			// })

			// mu.Lock()
			// async_page++
			// s.Prefix = "nhendl: doujin (#"+code+"): " + loading.Loading(pages, async_page) + " "
			// mu.Unlock()

		}(i)

		current_page++
	}

	wg.Wait()
	
	// increaseLoad(pages, limit_conc)
	
	// if(async_page == pages) {
	s.Suffix = " [async] (Completed)"
	// } else {
	// 	asyncProcessDownload(code, host)
	// }

}